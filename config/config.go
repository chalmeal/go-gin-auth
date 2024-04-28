package config

import (
	"fmt"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
)

func init() {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		fmt.Println("loade config err")
	}
}
