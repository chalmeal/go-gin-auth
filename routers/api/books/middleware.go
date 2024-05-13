/**
* books middleware
 */
package books

import (
	"go-gin-auth/common/connect"
	"go-gin-auth/common/sessions"
)

var (
	db   = connect.DbConnect()
	sess sessions.SessionInfo
)

func setUp() {
	db.AutoMigrate(&books{})
}
