/**
* auth info model
 */
package auths

import (
	"gorm.io/gorm"
)

// users expression to be authenticated
type users struct {
	UserId      string `gorm:"size:50;not null;unique" json:"user_id"`
	Name        string `gorm:"size:50;not null" json:"name"`
	Password    string `gorm:"size:500;not null" json:"password"`
	AccessToken string `gorm:"size:500" json:"access_token"`
	// timezone is UTC
	gorm.Model
}

func (u *users) getUserById(id string) *users {
	users := users{}
	db.Where("user_id=?", id).Find(&users)
	return &users
}

func (u *users) getPasswordById(id string) *users {
	users := users{}
	db.Select("user_id, password").Where("user_id=?", id).Find(&users)
	return &users
}

func (u *users) getAccessTokenById(id string) *users {
	users := users{}
	db.Select("access_token").Where("user_id=?", id).Find(&users)
	return &users
}

func (u *users) regUser(p users) error {
	err := db.Create(&p).Error
	if err != nil {
		return err
	}
	return nil
}

type getUserParam struct {
	UserId   string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type regUserParam struct {
	UserId   string `json:"userId" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
