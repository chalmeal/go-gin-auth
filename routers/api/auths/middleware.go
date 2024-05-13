/**
* auth middleware
 */
package auths

import (
	"errors"
	"fmt"
	"go-gin-auth/common/connect"
	"go-gin-auth/common/response"
	"go-gin-auth/common/sessions"
	"go-gin-auth/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	db   = connect.DbConnect()
	sess sessions.SessionInfo
)

func setUp() {
	db.AutoMigrate(&users{})
}

func VerifyAccessToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			sess.Clear(c)
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return []byte(config.GetInitKey("security", "JWT_SECRET_KEY")), nil
	})
	if err != nil {
		sess.Clear(c)
		response.Res(c, http.StatusInternalServerError, err)
		return
	}

	userId := t.Claims.(jwt.MapClaims)["user_id"]
	ut := rep.getAccessTokenById(userId.(string))
	if !t.Valid || ut.AccessToken != hashToken(t.Raw) {
		sess.Clear(c)
		response.Res(c, http.StatusUnauthorized, errors.New("unauthorized accessToken"))
		return
	}

	state := t.Claims.(jwt.MapClaims)["state"]
	if state != sess.GetState(c) {
		sess.Clear(c)
		response.Res(c, http.StatusUnauthorized, errors.New("csrf token verification failed"))
		return
	}
}
