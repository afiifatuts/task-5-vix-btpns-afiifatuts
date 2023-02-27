package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email"`
	Password string `gorm:"->;<-;notnull" json:"-"`
}
