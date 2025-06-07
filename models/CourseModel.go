package models

import (
	"errors"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Content         string        `gorm:"not null"`
	Room            Room          `gorm:"not null;foreignKey:RoomId;references:ID;onDelete:CASCADE"`
	CourseMachine   CourseMachine `gorm:"not null;foreignKey:CourseMachineId;references:ID;onDelete:CASCADE"`
	RoomId          uint
	CourseMachineId uint
}

func (c *Course) BeforeCreate(tx *gorm.DB) error {
	if err := checkExistence(tx, &Room{}, c.RoomId); err != nil {
		return errors.New("room not found")
	}

	if err := checkExistence(tx, &CourseMachine{}, c.CourseMachineId); err != nil {
		return errors.New("course machine not found")
	}

	return nil
}

func (c *Course) CreateCourse() (*Course, error) {
	if err := db.Create(&c).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return c, nil
}
