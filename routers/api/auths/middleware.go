/**
* auth middleware
 */
package auths

import (
	"errors"
	"fmt"
	"go-gin-auth/common/connect"
	"go-gin-auth/common/response"
	"go-gin-auth/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	db = connect.DbConnect()
)

func setUp() {
	db.AutoMigrate(&users{})
}

func VerifyAccessToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		secret, err := config.Cfg.Section("jwt").GetKey("JWT_SECRET_KEY")
		if err != nil {
			return nil, nil
		}
		return []byte(secret.String()), nil
	})
	if err != nil {
		log.Print("AccessToken verification failed.")
		response.Res(c, http.StatusInternalServerError, err)
		return
	}

	userId := t.Claims.(jwt.MapClaims)["user_id"]
	ut := rep.getAccessTokenById(userId.(string))
	if !t.Valid || ut.AccessToken != hashToken(t.Raw) {
		err = errors.New("unauthorized accessToken")
		log.Print(err)
		response.Res(c, http.StatusUnauthorized, err)
		return
	}
}
