package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Question          string          `gorm:"not null"`
	Answer            string          `gorm:"not null"`
	Hint1             string          `gorm:""`
	Hint2             string          `gorm:""`
	Hint3             string          `gorm:""`
	Point             uint            `gorm:"not null;default:1" validate:"required,min=1,max=255"`
	DifficultyLevel   DifficultyLevel `gorm:"not null;foreignKey:DifficultyLevelId;references:ID;onDelete:CASCADE" validate:"-"`
	Course            Course          `gorm:"not null;foreignKey:CourseId;references:ID;onDelete:CASCADE" validate:"-"`
	DifficultyLevelId uint
	CourseId          uint
}

func (q *Question) BeforeSave(tx *gorm.DB) error {
	if err := checkExistence(tx, &Course{}, q.CourseId); err != nil {
		return errors.New("course not found")
	}

	if err := checkExistence(tx, &DifficultyLevel{}, q.DifficultyLevelId); err != nil {
		return errors.New("diffculty level not found")
	}

	q.Answer = strings.ToLower(q.Answer)
	validate := validateInit()
	err := validate.Struct(q)
	if err != nil {
		return err
	}
	return nil
}

func (q *Question) CreateQuestion() (*Question, error) {
	if err := db.Create(&q).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return q, nil
}

func GetAllQuestions() ([]Question, error) {
	var questions []Question

	if err := db.Find(&questions).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return questions, nil
}
