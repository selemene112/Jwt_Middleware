package models

import (
	"jwt_token/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Admin struct {
	GormModel
	AdminID uint `gorm:"not null" json:"adminid"`
	// Admin *Admin
	FirstName string    `gorm:"not null" json:"first_name" validate:"required-First name is required"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email" validate:"required-Email is required,email-Invalid email format"`
	Password  string    `gorm:"not null" json:"password" validate:"required-Password is required,MinStringLength(6)-Password has to have a minimum length of 6 characters"`
	
}

func (u *Admin) BeforeCreate(tx *gorm.DB)(err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate 
		return 
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (p *Admin) BeforeCreateAdmin(tx *gorm.DB)(err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate 
		return 
	}

	err = nil
	return
}

func (p *Admin) BeforeUpdateAdmin(tx *gorm.DB)(err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate 
		return 
	}

	err = nil
	return
}
