/**
* books info model
 */
package books

import "gorm.io/gorm"

// book master expression
type books struct {
	BookId string `gorm:"size:50;not null;unique" json:"book_id"`
	Name   string `gorm:"size:50;not null" json:"name"`
	Author string `gorm:"size:50" json:"author"`
	// timezone is UTC
	gorm.Model
}

func (b *books) getAll() *[]books {
	books := []books{}
	db.Find(&books)
	return &books
}

type regBooksParam struct {
	BookId string `json:"bookId" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Author string `json:"author"`
}
