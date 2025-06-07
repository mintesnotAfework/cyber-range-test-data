package models

import "gorm.io/gorm"

type HackingMachine struct {
	gorm.Model
	ImageNameOrId string `gorm:"not null"`
}

func (hm *HackingMachine) CreateHackingMachine() error {
	if err := db.Create(&hm).Error; err != nil {
		return CustomErrorHandlerMiddleware(err)
	}

	return nil
}

func GetHackingMachine(id uint) (*HackingMachine, error) {
	var hm HackingMachine
	if err := db.First(&hm, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &hm, nil
}
