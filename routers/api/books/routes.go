/**
* books package entory point
 */
package books

import (
	"go-gin-auth/routers/api/auths"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {

	// execution middleware
	{
		setUp()
		r.Use(auths.VerifyAccessToken)
	}

	{
		// get all books
		r.GET("", getAllBooks)
		//register book
		r.POST("", registerBook)
	}

}
