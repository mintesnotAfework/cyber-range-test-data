package models

import (
	"strings"

	"gorm.io/gorm"
)

type DifficultyLevel struct {
	gorm.Model
	Level string `gorm:"not null;unique"`
}

type DifficultyLevelCreate struct {
	Level string `validate:"required,oneof='easy' 'medium' 'hard' 'very easy' 'insane'"`
}

func (d *DifficultyLevel) BeforeSave(tx *gorm.DB) error {
	d.Level = strings.ToLower(d.Level)

	dc := DifficultyLevelCreate{
		Level: d.Level,
	}
	validate := validateInit()
	err := validate.Struct(dc)
	if err != nil {
		return err
	}

	return nil
}

func (d *DifficultyLevel) CreateDifficultyLevel() (*DifficultyLevel, error) {
	if err := db.Create(&d).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return d, nil
}

func GetDifficultyLevelById(id uint) (*DifficultyLevel, error) {
	var difficultyLevel DifficultyLevel
	if err := db.First(&difficultyLevel, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &difficultyLevel, nil
}

func GetDifficultyLevelByLevel(level string) (*DifficultyLevel, error) {
	var difficultyLevel DifficultyLevel
	if err := db.Where("level = ?", level).First(&difficultyLevel).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &difficultyLevel, nil
}

func GetAllDifficultyLevels() ([]DifficultyLevel, error) {
	var difficultyLevels []DifficultyLevel
	if err := db.Find(&difficultyLevels).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return difficultyLevels, nil
}

func (d *DifficultyLevel) UpdateDifficultyLevel(id uint) (*DifficultyLevel, error) {
	var difficultyLevel DifficultyLevel
	if err := db.First(&difficultyLevel, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	if d.Level != "" {
		difficultyLevel.Level = d.Level
	}
	if err := db.Updates(&difficultyLevel).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &difficultyLevel, nil
}

func DeleteDifficultyLevel(id uint) error {
	subQuery := db.Table("diffculty_levels").
		Where("diffculty_levels.id = ?", id).
		Select("diffculty_levels.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&DifficultyLevel{}).Error; err != nil {
		return err
	}

	return nil
}
