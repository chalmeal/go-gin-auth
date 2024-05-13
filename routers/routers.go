/**
* entry point for REST API execution
 */
package routers

import (
	"go-gin-auth/config"
	"go-gin-auth/routers/api"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte(config.GetInitKey("security", "SESSION_KEY_PAIR")))
	router.Use(sessions.Sessions(config.GetInitKey("security", "SESSION_NAME"), store))
	{
		setUpRouter(router)
	}

	return router
}

func setUpRouter(router *gin.Engine) {
	a := router.Group("/api")
	{
		api.RegisterRouter(a)
	}
}
