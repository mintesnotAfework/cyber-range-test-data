package models

import (
	"errors"

	"gorm.io/gorm"
)

type QuestionStudent struct {
	gorm.Model
	Student    Student  `gorm:"not null;foreignKey:StudentId;onDelete:CASCADE"`
	Question   Question `gorm:"not null;foreignKey:QuestionId;onDelete:CASCADE"`
	StudentId  uint
	QuestionId uint
}

func (qs *QuestionStudent) BeforeSave(tx *gorm.DB) error {
	if err := checkExistence(tx, &Student{}, qs.StudentId); err != nil {
		return errors.New("student not found")
	}

	if err := checkExistence(tx, &Question{}, qs.QuestionId); err != nil {
		return errors.New("question not found")
	}
	return nil
}

func (qs *QuestionStudent) CreateQuestionStudent() (*QuestionStudent, error) {
	if err := db.Create(&qs).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return qs, nil
}

func GetAllQuestionStudent() ([]QuestionStudent, error) {
	var questionStudents []QuestionStudent
	if err := db.Find(&questionStudents).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return questionStudents, nil
}
