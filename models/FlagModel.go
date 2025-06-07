package models

import (
	"errors"

	"gorm.io/gorm"
)

type Flag struct {
	gorm.Model
	Value            string         `gorm:"not null"`
	MachineStudent   MachineStudent `gorm:"not null;foreignKey:MachineStudentId"`
	MachineStudentId uint           `gorm:"uniqueIndex:idx_MachineStudent_flag"`
}

func (f *Flag) BeforeCreate(tx *gorm.DB) error {
	if err := checkExistence(tx, &MachineStudent{}, f.MachineStudentId); err != nil {
		return errors.New("machine not found")
	}

	var counter int64
	if err := tx.Model(&Flag{}).Where("machine_student_id=?", f.MachineStudentId).Count(&counter).Error; err != nil {
		return errors.New("flag already exists")
	}
	if counter > 0 {
		return errors.New("flag already exists")
	}
	return nil
}

func (f *Flag) BeforeUpdate(tx *gorm.DB) error {
	return checkExistence(tx, &Machine{}, f.MachineStudentId)
}

func (f *Flag) CreateFlag() (*Flag, error) {
	if err := db.Create(&f).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return f, nil
}
