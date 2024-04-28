/**
* auth package entory point
 */
package auths

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {

	// execution middleware
	{
		setUp()
	}

	{
		// login
		r.POST("/login", login)
		// register user
		r.POST("", registerUser)
	}

}
