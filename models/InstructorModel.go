package models

import (
	"errors"

	"gorm.io/gorm"
)

type Instructor struct {
	gorm.Model
	AccountVerified bool `gorm:"not null;default:false"`
	User            User `gorm:"not null;foreignKey:UserId;references:ID;onDelete:CASCADE"`
	UserId          uint `gorm:"uniqueIndex:idx_instructor_user;"`
}

func (i *Instructor) BeforeCreate(tx *gorm.DB) error {
	if err := checkExistence(db, &User{}, i.UserId); err != nil {
		return errors.New("user not found")
	}

	return nil
}

func (i *Instructor) CreateInstructor() (*Instructor, error) {
	if err := db.Create(&i).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return i, nil
}

func GetAllInstructor() ([]Instructor, error) {
	var instructors []Instructor

	if err := db.Joins("JOIN users ON users.id = instructors.user_id").Find(&instructors).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return instructors, nil
}
