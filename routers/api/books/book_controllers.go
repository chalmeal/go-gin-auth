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
	b := books{
		BookId: p.BookId,
		Name:   p.Name,
		Author: p.Author,
	}
	err := rep.regBook(b)
	if err != nil {
		log.Printf("Register book error : %s", err)
		response.Res(c, http.StatusInternalServerError, err)
		return
	}

	r := "Register Book Success"

	response.Res(c, http.StatusOK, r)
}
