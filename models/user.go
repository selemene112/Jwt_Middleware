package models

import (
	"jwt_token/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)



type User struct{
GormModel
FullName string `gorm:"not null" json:"full_name" form:"full_name" validate:"required-Your full name is required"`//validate
Email string `gorm:"not null;uniqueIndex" json:"email" form:"email" validate:"required-Your email is required, email~invaild email format"`
Password string `gorm:"not null" json:"password" form:"password" validate:"required-Your password is required, midstringlength(6)~password has to have a minimum length of 6 characters"`
// Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
Products  []Product `json:"products"`

}

func (u *User) BeforeCreate(tx *gorm.DB)(err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate 
		return 
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}