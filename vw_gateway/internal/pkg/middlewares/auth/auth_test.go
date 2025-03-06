package auth_test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spewerspew/spew"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	_ "vw_user/internal/data"
	"vw_user/internal/data/dal/query"
	"vw_user/internal/pkg/middlewares/auth"
)

func dump(a ...interface{}) {
	spew.Dump(a...)
	fmt.Println()
}
func TestGenerateAndParseToken(t *testing.T) {
	jwtAuthorizer := &auth.JWTAuth{
		Secret:            "asf41a52fasd541fd5f1a",
		AccessExpireTime:  2 * time.Hour,
		RefreshExpireTime: 15 * 24 * time.Hour,
	}

	do := query.User
	user, err := do.First()
	require.NoError(t, err)

	atoken, rtoken, err := jwtAuthorizer.CreateToken(user.UserID, user.Username, user.IsAdmin)
	require.NoError(t, err)
	dump(atoken, rtoken)

	aclaims := new(auth.AccessTokenClaims)
	rclaims := new(auth.RefreshTokenClaims)
	at, err := jwt.ParseWithClaims(atoken, aclaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtAuthorizer.Secret), nil
	})
	require.NoError(t, err)
	rt, err := jwt.ParseWithClaims(rtoken, rclaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtAuthorizer.Secret), nil
	})
	require.NoError(t, err)

	dump(at.Claims)
	dump(rt.Claims)
}
