package models

import (
	"errors"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	IsPremiem bool `gorm:"not null;default:false"`
	User      User `gorm:"not null;foreignKey:UserId;references:ID;onDelete:CASCADE"`
	UserId    uint `gorm:"uniqueIndex:idx_student_user;"`
}

func (s *Student) BeforeSave(tx *gorm.DB) error {
	if err := checkExistence(tx, &User{}, s.UserId); err != nil {
		return errors.New("user not found")
	}
	return nil
}

func (s *Student) CreateStudent() (*Student, error) {
	if err := db.Create(&s).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return s, nil
}

func GetAllStudents() ([]Student, error) {
	var students []Student

	if err := db.Find(&students).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return students, nil
}
