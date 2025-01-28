package auth

import (
	"context"
	kerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
	"vw_user/internal/conf"
	"vw_user/internal/pkg/ecode/errdef"
)

type JWTAuth struct {
	Secret            string
	AccessExpireTime  time.Duration
	RefreshExpireTime time.Duration
}

func NewJWTAuth(cfg *conf.JWT) *JWTAuth {
	return &JWTAuth{
		Secret:            cfg.Secret,
		AccessExpireTime:  time.Duration(cfg.AccessTokenExpiration) * time.Hour,
		RefreshExpireTime: time.Duration(cfg.RefreshTokenExpiration) * 24 * time.Hour,
	}
}

const (
	BearerTokenPrefix  = "Bearer"
	AccessTokenHeader  = "Authorization"
	RefreshTokenHeader = "Refresh-Token"

	// TokenIssuer is the default token issuer
	TokenIssuer = "ZEY_HUNTER_ETL"
)

// AccessTokenClaims access token claims, which is used to verify access token
type AccessTokenClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func NewAccessTokenClaims() *AccessTokenClaims {
	return &AccessTokenClaims{}
}
func (t *AccessTokenClaims) isExpired() bool {
	// if token.ExpiresAt before now, the token is expired, return true
	//else, return false
	return t.ExpiresAt.Time.Before(time.Now())
}

// padding pads the access token claims
func (t *AccessTokenClaims) padding(userID int64, username string, isAdmin bool, duration time.Duration) {
	t.UserID = userID
	t.Username = username
	t.IsAdmin = isAdmin
	t.Issuer = TokenIssuer
	t.ExpiresAt = jwt.NewNumericDate(time.Now().Add(duration))

}

// getTokenString returns the token string of access token claims,which is padded
func (t *AccessTokenClaims) getTokenString(secret string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", errdef.ErrCreateTokenFailed
	}
	return
}

// RefreshTokenClaims refresh token claims, which is used to refresh access token
type RefreshTokenClaims struct {
	jwt.RegisteredClaims
}

func NewRefreshTokenClaims() *RefreshTokenClaims {
	return &RefreshTokenClaims{}
}
func (t *RefreshTokenClaims) isExpired() bool {
	return t.ExpiresAt.Time.Before(time.Now())
}

// padding pads the refresh token claims
func (t *RefreshTokenClaims) padding(duration time.Duration) {
	t.Issuer = TokenIssuer
	t.ExpiresAt = jwt.NewNumericDate(time.Now().Add(duration))
}

// getTokenString returns the token string of refresh token claims,which is padded
func (t *RefreshTokenClaims) getTokenString(secret string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", errdef.ErrCreateTokenFailed
	}
	return
}

// CreateToken create access token and refresh token
func (jwtAuthorizer *JWTAuth) CreateToken(userID int64, username string, isAdmin bool) (accessToken, refreshToken string, err error) {
	// Create access token
	atokenClaims := NewAccessTokenClaims()
	atokenClaims.padding(userID, username, isAdmin, jwtAuthorizer.AccessExpireTime)
	accessToken, err = atokenClaims.getTokenString(jwtAuthorizer.Secret)
	if err != nil {
		return "", "", err
	}

	// Create refresh token
	rtokenClaims := NewRefreshTokenClaims()
	rtokenClaims.padding(jwtAuthorizer.RefreshExpireTime)
	refreshToken, err = rtokenClaims.getTokenString(jwtAuthorizer.Secret)
	if err != nil {
		return "", "", err
	}
	return
}

func JwtAuth(secret string, accessTokenExpireTime time.Duration) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {

				// Get access token from header
				atoken, err := getToken(tr, secret, AccessTokenHeader)
				if err != nil {
					return nil, err
				}
				// Get refresh token from header
				rtoken, err := getToken(tr, secret, RefreshTokenHeader)
				if err != nil {
					return nil, err
				}

				aclaims, tok := atoken.Claims.(*AccessTokenClaims)
				rclaims, tok := rtoken.Claims.(*RefreshTokenClaims)
				if tok && atoken.Valid {
					//判断是否过期，是否需要刷新
					if aclaims.isExpired() {
						if rclaims.isExpired() { // refresh token expired，need login again
							return nil, errdef.ErrRefreshTokenExpired
						} else {
							// refresh access token
							atokenClaims := NewAccessTokenClaims()
							atokenClaims.padding(aclaims.UserID, aclaims.Username, aclaims.IsAdmin, accessTokenExpireTime)
							accessToken, err := atokenClaims.getTokenString(secret)
							if err != nil {
								return nil, err
							}
							return nil, (errdef.ErrAccessTokenExpired.(*kerr.Error)).WithMetadata(map[string]string{
								"access_token": accessToken,
							})
						}
					}
				} else {
					return nil, errdef.ErrTokenInvalid
				}
			}
			return handler(ctx, req)
		}
	}
}
func getToken(tr transport.Transporter, secret, headerString string) (tokenClaims *jwt.Token, err error) {
	tokenString := tr.RequestHeader().Get(headerString) //get token from header,which contains "Bearer "
	auths := strings.SplitN(tokenString, " ", 2)        //get raw token
	if len(auths) != 2 || !strings.EqualFold(auths[0], BearerTokenPrefix) {
		return nil, errdef.ErrTokenMissing
	}
	if headerString == AccessTokenHeader {
		atokenClaims := new(AccessTokenClaims)
		tokenClaims, err = jwt.ParseWithClaims(auths[1], atokenClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			return nil, errdef.ErrParseTokenFailed
		}
	} else { //headerString == RefreshTokenHeader
		rtokenClaims := new(jwt.RegisteredClaims)
		tokenClaims, err = jwt.ParseWithClaims(auths[1], rtokenClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			return nil, errdef.ErrParseTokenFailed
		}
	}
	return
}
