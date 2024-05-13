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
	cnf := config.GetInitKeyFatal("database", "DB_PATH")
	db, err := gorm.Open(mysql.Open(cnf), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
