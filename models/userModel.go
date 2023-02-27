package models

import (
	"time"

	"github.com/afiifatuts/go-authentication/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"not null" json:"username" form:"username" valid:"required~ Username is required"`
	Email     string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required, email~Invalid email format"`
	Password  string     `gorm:"->;<-;notnull" json:"-"`
	Photos    []Photo    `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"-"`
	CreatedAt *time.Time `json:"-,omitempty"`
	UpdatedAt *time.Time `json:"-,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
