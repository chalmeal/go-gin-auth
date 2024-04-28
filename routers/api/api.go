/**
* api entory point
 */
package api

import (
	"go-gin-auth/routers/api/auths"
	"go-gin-auth/routers/api/books"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup) {
	// Auth
	auths.RegisterRouter(router.Group("/auth"))
	// Book
	books.RegisterRouter(router.Group("/book"))
}
