package models

import (
	"strings"

	"gorm.io/gorm"
)

type OperatingSystemType struct {
	gorm.Model
	Type string `gorm:"not null;unique;" validate:"required,oneof='windows' 'macos' 'linux' 'android' 'other' 'ios'"`
}

func (o *OperatingSystemType) BeforeSave(tx *gorm.DB) error {
	o.Type = strings.ToLower(o.Type)

	validate := validateInit()
	err := validate.Struct(o)
	if err != nil {
		return err
	}
	return nil
}

func (o *OperatingSystemType) CreateOperatingSystemType() (*OperatingSystemType, error) {
	if err := db.Create(&o).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return o, nil
}

func GetOperatingSystemTypeById(id uint) (*OperatingSystemType, error) {
	var operatingSystemType OperatingSystemType
	if err := db.First(&operatingSystemType, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &operatingSystemType, nil
}

func GetOperatingSystemTypeByType(osType string) (*OperatingSystemType, error) {
	var operatingSystemType OperatingSystemType
	if err := db.Where("type = ?", osType).First(&operatingSystemType).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &operatingSystemType, nil
}

func GetAllOperatingSystemType() ([]OperatingSystemType, error) {
	var operatingSystemTypes []OperatingSystemType
	if err := db.Find(&operatingSystemTypes).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return operatingSystemTypes, nil
}

func (o *OperatingSystemType) UpdateOperatingSystemType(id uint) (*OperatingSystemType, error) {
	var operatingSystemType OperatingSystemType
	if err := db.First(&operatingSystemType, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	if o.Type != "" {
		operatingSystemType.Type = o.Type
	}
	if err := db.Updates(&operatingSystemType).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &operatingSystemType, nil
}

func DeleteOperatingSystemType(id uint) error {
	subQuery := db.Table("operating_system_types").
		Where("operating_system_types.id = ?", id).
		Select("operating_system_types.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&OperatingSystemType{}).Error; err != nil {
		return err
	}

	return nil
}
