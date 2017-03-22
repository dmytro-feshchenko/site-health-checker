package models

import (
	"github.com/jinzhu/gorm"
)

// User - model for user
type User struct {
	Model

	FirstName string `json:"first_name" gorm:"size:255" form:"first_name" query:"first_name"`
	LastName  string `json:"last_name" gorm:"size:255" form:"last_name" query:"last_name"`
	FullName  string `json:"full_name" gorm:"size:510"`
	Username  string `json:"username" gorm:"size:255" form:"username" query:"username"`
	Password  []byte `json:"-" gorm:"type:bytea"`

	// Birthday        time.Time  `json:"birthday" form:"birthday" query:"birthday"`
	// CreditCard      CreditCard `json:"credit_card"` // One-To-One relationship
	Email           string `json:"email" form:"email" query:"email"`
	IsEmailVerified bool   `sql:"DEFAULT:false" json:"is_email_verified"`
	IsSubscribed    bool   `sql:"DEFAULT:true" json:"is_subscribed"`
	// Email      []Email    `json:"emails"`      // One-To-Many relationship
}

// Email - email for user
type Email struct {
	Model
	UserID     uint   `json:"user_id",gorm:"index"` // foreign key (belongs to)
	Email      string `json:"email",gorm:"type:varchar(254);unique_index"`
	Subscribed bool   `json:"subscribed"`
}

// // CreditCard - credit card info
// type CreditCard struct {
// 	Model
// 	UserID uint   `json:"user_id"`
// 	Number string `json:"number"`
// }

//BeforeSave - before saving user
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// set full name field if first_name or last_name have been changed
	fullName := u.FirstName + " " + u.LastName
	if u.FullName != fullName {
		tx.Model(u).Update("FullName", fullName)
	}
	return
}
