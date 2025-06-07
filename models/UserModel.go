package models

import (
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName      string `gorm:"unique;not null;"`
	FirstName     string `gorm:"not null;"`
	LastName      string `gorm:"not null;"`
	MiddleName    string `gorm:""`
	Email         string `gorm:"unique;not null;"`
	Phone         string `gorm:"unique;not null;"`
	Password      string `gorm:"not null"`
	EmailVerified bool   `gorm:"not null;default:false"`
	PhoneVerified bool   `gorm:"not null;default:false"`
	Locked        bool   `gorm:"not null:default:false"`
}

type UserCreate struct {
	UserName   string `validate:"required,username"`
	FirstName  string `validate:"required,alpha,min=1,max=15"`
	LastName   string `validate:"required,alpha,min=1,max=15"`
	MiddleName string `validate:"omitempty,alpha,min=1,max=15"`
	Email      string `validate:"required,email"`
	Phone      string `validate:"required,phone"`
	Password   string `validate:"required"`
}

type UserUpdate struct {
	UserName   string `validate:"omitempty,username"`
	FirstName  string `validate:"omitempty,alpha,min=1,max=15"`
	LastName   string `validate:"omitempty,alpha,min=1,max=15"`
	MiddleName string `validate:"omitempty,alpha,min=1,max=15"`
	Email      string `validate:"omitempty,email"`
	Phone      string `validate:"omitempty,phone"`
	Password   string `validate:"omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UserName = strings.ToLower(u.UserName)
	u.Email = strings.ToLower(u.Email)

	uc := UserCreate{
		UserName:   u.UserName,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Email:      u.Email,
		Phone:      u.Phone,
		Password:   u.Password,
	}

	validate := validateInit()
	err := validate.Struct(uc)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UserName = strings.ToLower(u.UserName)
	u.Email = strings.ToLower(u.Email)

	uu := UserUpdate{
		UserName:   u.UserName,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Email:      u.Email,
		Phone:      u.Phone,
		Password:   u.Password,
	}

	validate := validateInit()
	err := validate.Struct(uu)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) CreateUser() (*User, error) {
	if err := db.Create(&u).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return u, nil
}

func GetAllUsers() ([]User, error) {
	var users []User

	if err := db.Find(&users).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return users, nil
}
