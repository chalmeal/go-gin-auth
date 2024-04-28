/**
* auth controller
 */
package auths

import (
	"errors"
	"go-gin-auth/common/response"
	"go-gin-auth/config"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var rep users

// claim paramter
type ClaimsRecord struct {
	UserId   string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
	jwt.RegisteredClaims
}

/*
* users login authentication
*
* @param: UserId string
* @param: Password string
 */
func login(c *gin.Context) {
	p := new(getUserParam)
	if err := c.Bind(&p); err != nil {
		log.Printf("Validation error: %s", err)
		response.Res(c, http.StatusBadRequest, err)
		return
	}

	users := rep.getPasswordById(p.UserId)
	if users.Password == "" || hashPW(p.Password) != users.Password {
		response.Res(c, http.StatusUnauthorized, errors.New("unauthorized UserID or Password"))
		return
	}

	claims := &ClaimsRecord{
		users.UserId,
		users.Password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := config.Cfg.Section("jwt").GetKey("JWT_SECRET_KEY")
	if err != nil {
		log.Printf("Login error : %s", err)
		response.Res(c, http.StatusInternalServerError, err)
		return
	}
	t, err := token.SignedString([]byte(secret.String()))
	if err != nil {
		log.Printf("Login error : %s", err)
		response.Res(c, http.StatusInternalServerError, err)
		return
	}

	if err = updateAccessToken(users.UserId, t); err != nil {
		response.Res(c, http.StatusInternalServerError, err)
		return
	}

	r := map[string]string{"AccessToken": t}

	response.Res(c, http.StatusOK, r)
}

/*
* register users
*
* @param: UserId string
* @param: Name string
* @param: Password string
 */
func registerUser(c *gin.Context) {
	p := new(regUserParam)
	if err := c.Bind(&p); err != nil {
		log.Printf("Validation error: %s", err)
		response.Res(c, http.StatusBadRequest, err)
		return
	}
	users := users{
		UserId:   p.UserId,
		Name:     p.Name,
		Password: hashPW(p.Password),
	}
	result := db.Create(&users)
	if result.Error != nil {
		log.Printf("Register users error : %s", result.Error)
		response.Res(c, http.StatusInternalServerError, result.Error)
		return
	}

	r := "Register User Success"

	response.Res(c, http.StatusOK, r)
}

/*
* update accessToken
*
* @param: id string
* @param: t string
*
* @return: error
 */
func updateAccessToken(id string, t string) error {
	if u := rep.getUserById(id); u.UserId == "" {
		log.Printf("Not found users : %s", id)
		return errors.New("not found users")
	}

	result := db.Model(&users{}).Where("user_id=?", id).Update("access_token", hashToken(t))
	if result.Error != nil {
		log.Printf("Update accessToken error : %s", result.Error)
		return errors.New("update accessToken error")
	}
	return nil
}
