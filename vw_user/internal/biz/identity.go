package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"unicode/utf8"
	"util/snowflake"
	"vw_user/internal/data/dal/model"
	"vw_user/internal/data/dal/query"
	"vw_user/internal/pkg/captcha"
	"vw_user/internal/pkg/ecode/errdef"
	"vw_user/internal/pkg/helper"
	"vw_user/internal/pkg/middlewares/auth"
)

type UserIdentityRepo interface {
	// AddExpForLogin add user exp for login
	AddExpForLogin(ctx context.Context, userId int64) error

	// SetCodeCache  set <email,code> in redis
	SetCodeCache(ctx context.Context, email, code string) error

	// DeleteCodeFromCache delete <email,code> from redis
	DeleteCodeFromCache(ctx context.Context, email string)

	// GetCodeFromCache get verify-code from redis
	GetCodeFromCache(ctx context.Context, email string) (code string, err error)
}

type RegisterInfo struct {
	Username       string
	Password       string
	RepeatPassword string
	Gender         int32
	Code           string
	Email          string
	Signature      string
	Birthday       time.Time
}
type UserIdentityUsecase struct {
	infoRepo      UserInfoRepo
	identityRepo  UserIdentityRepo
	jwtAuthorizer *auth.JWTAuth
	email         *captcha.Email
	logger        *log.Helper
}

func NewUserIdentityUsecase(idRepo UserIdentityRepo, infoRepo UserInfoRepo, jwt *auth.JWTAuth, ecfg *captcha.Email, logger log.Logger) *UserIdentityUsecase {
	return &UserIdentityUsecase{
		identityRepo:  idRepo,
		infoRepo:      infoRepo,
		jwtAuthorizer: jwt,
		email:         ecfg,
		logger:        log.NewHelper(logger),
	}
}
func (uc *UserIdentityUsecase) Register(ctx context.Context, req *http.Request, registerInfo *RegisterInfo) (atoken, rtoken string, err error) {
	newUser := &model.User{
		UserID:    snowflake.GetID(),
		Username:  registerInfo.Username,
		Password:  registerInfo.Password,
		Gender:    int(registerInfo.Gender),
		Email:     registerInfo.Email,
		Signature: registerInfo.Signature,
		Birthday:  registerInfo.Birthday,
	}
	// check register info
	err = uc.checkRegisterInfo(newUser, registerInfo.RepeatPassword)
	if err != nil {
		return "", "", err
	}

	//check code captcha
	code, err := uc.identityRepo.GetCodeFromCache(ctx, registerInfo.Email)
	if errors.Is(err, redis.Nil) {
		return "", "", errdef.ErrVerifyCodeExpired
	}
	if code != registerInfo.Code {
		return "", "", errdef.ErrVerifyCodeNotMatch
	}

	// encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", errdef.ErrEncryptPasswordFailed
	}
	newUser.Password = string(hashedPassword)

	// set user default avatar
	//// create user dir
	userDir := filepath.Join(resourcePath, strconv.FormatInt(newUser.UserID, 10))
	err = helper.CreateDir(userDir, os.ModePerm)
	defer func() {
		/* this defer is used to remove the user dir
		if there is any error Behind this logic*/
		if err != nil {
			_ = os.RemoveAll(userDir)
		}
	}()
	if err != nil {
		return "", "", errdef.ErrCreateUserDirFailed
	}
	avatar, err := helper.FormFile(req, "avatar")
	if err != nil {
		return "", "", errdef.ErrReadAvatarFailed
	}
	var avatarFilePath string
	if avatar == nil {
		// user don't upload avatar, set default avatar
		avatarFilePath = filepath.Join(userDir, "avatar.jpg")
		defaultAvatar, err := helper.ReadFileContent(defaultAvatarPath)
		if err != nil {
			return "", "", errdef.ErrReadAvatarFailed
		}
		err = os.WriteFile(avatarFilePath, defaultAvatar, os.ModePerm)
		if err != nil {
			return "", "", errdef.ErrSaveAvatarFailed
		}
	} else {
		avatarFilePath = filepath.Join(userDir, "avatar"+filepath.Ext(avatar.Filename))
		err = helper.WriteToNewFile(avatar, avatarFilePath)
		if err != nil {
			return "", "", errdef.ErrSaveAvatarFailed
		}
	}
	newUser.AvatarPath = avatarFilePath

	//create user's default follow_list and favorites
	defaultFavorites := &model.UserFavorite{
		FavoriteID:   snowflake.GetID(),
		UserID:       newUser.UserID,
		FavoriteName: "默认收藏夹",
		Description:  "",
		IsPrivate:    1,
	}
	privateFavorites := &model.UserFavorite{
		FavoriteID:   snowflake.GetID(),
		UserID:       newUser.UserID,
		FavoriteName: "私密收藏夹",
		Description:  "",
		IsPrivate:    -1,
	}
	userLevel := &model.UserLevel{
		UserID: newUser.UserID,
	}
	defaultFollowList := &model.FollowList{
		ListID:   snowflake.GetID(),
		UserID:   newUser.UserID,
		ListName: "默认关注列表",
	}

	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = tx.User.Create(newUser)
	if err != nil {
		return "", "", err
	}

	err = tx.FollowList.Create(defaultFollowList)
	if err != nil {
		return "", "", err
	}

	err = tx.UserFavorite.Create(defaultFavorites, privateFavorites)
	if err != nil {
		return "", "", err
	}

	err = tx.UserLevel.Create(userLevel)
	if err != nil {
		return "", "", err
	}
	tx.Commit()
	//
	atoken, rtoken, jwtErr := uc.jwtAuthorizer.CreateToken(newUser.UserID, newUser.Username, newUser.IsAdmin)
	if jwtErr != nil {
		return "", "", errdef.ErrCreateTokenFailed
	}
	return atoken, rtoken, nil
}

func (uc *UserIdentityUsecase) Login(ctx context.Context, name string, password string) (accessToken, refreshToken string, err error) {
	userInfo, err := uc.infoRepo.GetUserSummaryInfoByUsername(ctx, name)
	if err != nil {
		return "", "", err
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password)); err != nil {
		return "", "", errdef.ErrUserPasswordError
	}

	//login successfully,then get token
	accessToken, refreshToken, err = uc.jwtAuthorizer.CreateToken(userInfo.UserID, userInfo.Username, userInfo.IsAdmin)
	if err != nil {
		return "", "", errdef.ErrCreateTokenFailed
	}

	err = uc.identityRepo.AddExpForLogin(ctx, userInfo.UserID)
	if err != nil {
		//TODO:若用户经验更新失败，使用异步方法将消息传送到消息队列中，后台异步更新用户经验，现阶段不做处理
		return
	}
	//
	//// update or create user cache in redis
	////TODO: 缓存用户信息到redis中,如果原本存在缓存，则不做处理，如果更新出错，则由后台异步更新缓存
	//uc.identityRepo.SetUserCache(ctx, userInfo.UserID)
	return
}

// SendCodeCaptcha send code captcha to user's email
func (uc *UserIdentityUsecase) SendCodeCaptcha(ctx context.Context, email string) (code string, err error) {
	// when user get code captcha, delete the old code in redis first
	uc.identityRepo.DeleteCodeFromCache(ctx, email)
	code = uc.email.CreateVerificationCode()
	err = uc.email.SendCode(ctx, email, code)
	if err != nil {
		return "", err
	}
	err = uc.identityRepo.SetCodeCache(ctx, email, code)
	if err != nil {
		return "", err
	}
	return
}

// checkRegisterInfo check the register info is valid or not
func (uc *UserIdentityUsecase) checkRegisterInfo(checkInfo *model.User, repeatPassword string) error {
	userdo := query.User
	count, _ := userdo.Where(userdo.Username.Eq(checkInfo.Username)).Count()

	switch {
	case count > 0: //已有同名用户
		return errdef.ErrUserAlreadyExist
	case len(checkInfo.Password) < 6: //密码长度小于6位
		return errdef.ErrPasswordTooShort
	case checkInfo.Password != repeatPassword: //第一次输入的密码与第二次输入的密码不一致.
		return errdef.ErrPasswordNotMatch
	case utf8.RuneCountInString(checkInfo.Signature) > 25:
		return errdef.ErrSignatureTooLong
	}
	return nil
}
