/**
* response expressions to be returned
 */
package response

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// return response value
func Res(c *gin.Context, code int, r interface{}) {
	defer c.AbortWithStatus(code)
	if code != http.StatusOK {
		log.Printf("Failure API status code: %v\n", code)
		c.JSON(code, gin.H{
			"Status":  code,
			"Message": fmt.Sprint(r),
		})
	} else {
		c.JSON(code, gin.H{
			"Status": code,
			"Result": r,
		})
	}
}
