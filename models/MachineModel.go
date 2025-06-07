package models

import (
	"errors"

	"gorm.io/gorm"
)

type Machine struct {
	gorm.Model
	ImageNameOrID         string              `gorm:"not null"`
	Room                  Room                `gorm:"not null;foreignKey:RoomId;references:ID;onDelete:CASCADE"`
	CourseMachine         CourseMachine       `gorm:"not null;foreignKey:CourseMachineId;references:ID;onDelete:CASCADE"`
	OperatingSystemType   OperatingSystemType `gorm:"not null;foreignKey:OperatingSystemTypeId;references:ID;onDelete:CASCADE"`
	RoomId                uint
	CourseMachineId       uint
	OperatingSystemTypeId uint
}

func (m *Machine) BeforeCreate(tx *gorm.DB) error {
	if err := checkExistence(tx, &OperatingSystemType{}, m.OperatingSystemTypeId); err != nil {
		return errors.New("operating system type not found")
	}

	if err := checkExistence(tx, &Room{}, m.RoomId); err != nil {
		return errors.New("room not found")
	}

	if err := checkExistence(tx, &CourseMachine{}, m.CourseMachineId); err != nil {
		return errors.New("course machine not found")
	}
	return nil
}

func (m *Machine) CreateMachine() (*Machine, error) {
	if err := db.Create(&m).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return m, nil
}

func GetAllMachine() ([]Machine, error) {
	var machines []Machine

	if err := db.Find(&machines).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return machines, nil
}
