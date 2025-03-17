package errdef

import (
	kerr "github.com/go-kratos/kratos/v2/errors"
	"vw_user/internal/pkg/ecode"
)

func init() {
	ErrFollowOtherUserFailed = kerr.New(ecode.FOLLOW_FollowOtherUserFailed, "关注用户失败", "关注用户失败，请稍后再试")
	ErrUnfollowOtherUserFailed = kerr.New(ecode.FOLLOW_UnfollowOtherUserFailed, "取消关注用户失败", "取消关注用户失败，请稍后再试")
}

var (
	ErrFollowOtherUserFailed   *kerr.Error
	ErrUnfollowOtherUserFailed *kerr.Error
)
