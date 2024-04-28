/**
* books middleware
 */
package books

import (
	"go-gin-auth/common/connect"
)

var (
	db = connect.DbConnect()
)

func setUp() {
	db.AutoMigrate(&books{})
}
