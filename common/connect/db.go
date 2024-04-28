/*
* external connection information definition
 */
package connect

import (
	"go-gin-auth/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// db connect
func DbConnect() *gorm.DB {
	cnf, err := config.Cfg.Section("database").GetKey("DB_PATH")
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(mysql.Open(cnf.String()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
