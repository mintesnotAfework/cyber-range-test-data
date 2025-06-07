package models

import (
	"errors"

	"gorm.io/gorm"
)

type CourseStudent struct {
	gorm.Model
	Done      bool    `gorm:"not null;default:false"`
	Student   Student `gorm:"not null;foreignKey:StudentId;onDelete:CASCADE"`
	Course    Course  `gorm:"not null;foreignKey:CourseId;onDelete:CASCADE"`
	StudentId uint
	CourseId  uint
}

func (cs *CourseStudent) BeforeCreate(tx *gorm.DB) error {
	if err := checkExistence(tx, &Student{}, cs.StudentId); err != nil {
		return errors.New("student not found")
	}

	if err := checkExistence(tx, &Course{}, cs.CourseId); err != nil {
		return errors.New("course not found")
	}
	return nil
}

func (cs *CourseStudent) CreateCourseStudent() (*CourseStudent, error) {
	if err := db.Create(&cs).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return cs, nil
}

func GetAllCourseStudent() ([]CourseStudent, error) {
	var courseStudents []CourseStudent

	if err := db.Find(&courseStudents).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return courseStudents, nil
}
