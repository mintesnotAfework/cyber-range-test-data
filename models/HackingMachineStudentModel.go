package models

import "gorm.io/gorm"

type HackingMachineStudent struct {
	gorm.Model
	ContainerNameorId string         `gorm:"not null"`
	Student          Student        `gorm:"not null;foreignKey:StudentId;references:ID;onDelete:CASCADE"`
	HackingMachine   HackingMachine `gorm:"not null;foreignKey:HackingMachineId;references:ID;onDelete:CASCADE"`
	StudentId        uint
	HackingMachineId uint
}
