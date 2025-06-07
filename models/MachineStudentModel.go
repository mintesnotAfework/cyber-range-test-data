package models

import (
	"errors"

	"gorm.io/gorm"
)

type MachineStudent struct {
	gorm.Model
	ContainerNameOrID string  `gorm:"not null"`
	Student           Student `gorm:"not null;foreignKey:StudentId;onDelete:CASCADE"`
	Machine           Machine `gorm:"not null;foreignKey:MachineId;onDelete:CASCADE"`
	Done              bool    `gorm:"not null;default:false;"`
	StudentId         uint
	MachineId         uint
}

func (ms *MachineStudent) BeforeCreate(tx *gorm.DB) error {
	if err := checkExistence(tx, &Student{}, ms.StudentId); err != nil {
		return errors.New("student not found")
	}

	if err := checkExistence(tx, &Machine{}, ms.MachineId); err != nil {
		return errors.New("machine not found")
	}
	return nil
}

func (ms *MachineStudent) CreateMachineStudent() (*MachineStudent, error) {
	if err := db.Create(&ms).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return ms, nil
}

func GetAllMachineStudent() ([]MachineStudent, error) {
	var machineStudents []MachineStudent

	if err := db.Find(&machineStudents).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return machineStudents, nil
}
