package main

import (
	"go-gin-auth/config"
	"go-gin-auth/routers"
)

func main() {
	port := config.GetInitKeyFatal("server", "SERVER_PORT")

	r := routers.InitRouter()
	r.Run(":" + port)
}
