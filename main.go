package main

import (
	"go-gin-auth/config"
	"go-gin-auth/routers"
	"log"
)

func main() {
	port, err := config.Cfg.Section("server").GetKey("SERVER_PORT")
	if err != nil {
		log.Fatal(err)
	}

	r := routers.InitRouter()
	r.Run(":" + port.String())
}
