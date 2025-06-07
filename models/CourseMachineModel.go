package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type CourseMachine struct {
	gorm.Model
	Title             string          `gorm:"unique;not null;"`
	Description       string          `gorm:"not null;"`
	Point             uint            `gorm:"not null"`
	DifficultyLevel   DifficultyLevel `gorm:"not null;foreignKey:DifficultyLevelId;references:ID;onDelete:CASCADE"`
	DifficultyLevelId uint
}

type CourseMachineCreate struct {
	Title       string `validate:"required,min=2,max=50"`
	Description string `validate:"required,min=10,max=200"`
	Point       uint   `validate:"required,min=1,max=255"`
}

type CourseMachineUpdate struct {
	Title       string `validate:"omitempty,min=2,max=50"`
	Description string `validate:"omitempty,min=10,max=200"`
	Point       uint   `validate:"omitempty,min=1,max=255"`
}

func (cm *CourseMachine) BeforeSave(tx *gorm.DB) error {
	cm.Title = strings.ToLower(cm.Title)
	return nil
}
func (cm *CourseMachine) BeforeCreate(tx *gorm.DB) error {
	if err := checkExistence(tx, &DifficultyLevel{}, cm.DifficultyLevelId); err != nil {
		return errors.New("diffculty level not found")
	}

	cmc := CourseMachineCreate{
		Title:       cm.Title,
		Description: cm.Description,
		Point:       cm.Point,
	}

	validate := validateInit()
	err := validate.Struct(cmc)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CourseMachine) BeforeUpdate(tx *gorm.DB) error {
	cmu := CourseMachineUpdate{
		Title:       cm.Title,
		Description: cm.Description,
		Point:       cm.Point,
	}

	validate := validateInit()
	err := validate.Struct(cmu)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CourseMachine) CreateCourseMachine() (*CourseMachine, error) {
	if err := db.Create(&cm).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return cm, nil
}

func GetCourseMachineById(id uint) (*CourseMachine, error) {
	var courseMachine CourseMachine

	if err := db.First(&courseMachine, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &courseMachine, nil
}

func GetCourseMachineByCourseId(id uint) (*CourseMachine, error) {
	var courseMachine CourseMachine

	if err := db.Joins("JOIN course_machines ON course_machines.id=courses.course_machine_id").
		Where("courses.id=?", id).First(&courseMachine).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &courseMachine, nil
}

func (cm *CourseMachine) UpdateCourseMachineById(id uint) (*CourseMachine, error) {
	var courseMachine CourseMachine
	if err := db.First(&courseMachine, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	if cm.Title != "" && cm.Title != courseMachine.Title {
		courseMachine.Title = cm.Title
	}
	if cm.Description != "" && cm.Description != courseMachine.Description {
		courseMachine.Description = cm.Description
	}
	if cm.Point != 0 && cm.Point != courseMachine.Point {
		courseMachine.Point = cm.Point
	}
	if cm.DifficultyLevelId != 0 && cm.DifficultyLevelId != courseMachine.DifficultyLevelId {
		courseMachine.DifficultyLevelId = cm.DifficultyLevelId
	}

	if err := db.Where("ID=?", id).Updates(&courseMachine).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &courseMachine, nil
}

func DeleteCourseMachineById(id uint) (*CourseMachine, error) {
	subQuery := db.Table("course_machines").
		Where("course_machines.id = ?", id).
		Select("course_machines.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&CourseMachine{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &CourseMachine{}, nil
}
