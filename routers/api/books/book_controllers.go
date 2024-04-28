/**
* books controller
 */
package books

import (
	"go-gin-auth/common/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var rep books

/*
* get all books
*
* @return: []books
 */
func getAllBooks(c *gin.Context) {
	b := rep.getAll()
	response.Res(c, http.StatusOK, b)
}

/*
* register book
*
* @param: BookId string
* @param: Name string
* @param: Author string
 */
func registerBook(c *gin.Context) {
	p := new(regBooksParam)
	if err := c.Bind(&p); err != nil {
		log.Printf("Validation error: %s", err)
		response.Res(c, http.StatusBadRequest, err)
		return
	}
	book := books{
		BookId: p.BookId,
		Name:   p.Name,
		Author: p.Author,
	}
	result := db.Create(&book)
	if result.Error != nil {
		log.Printf("Register book error : %s", result.Error)
		response.Res(c, http.StatusInternalServerError, result.Error)
		return
	}

	r := "Register Book Success"

	response.Res(c, http.StatusOK, r)
}
