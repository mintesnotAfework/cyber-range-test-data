package models

import (
	"errors"

	"gorm.io/gorm"
)

type RevokedTokenInstructor struct {
	gorm.Model
	Instructor     Instructor   `gorm:"not null;foreignKey:InstructorId;references:ID;onDelete:CASCADE"`
	RevokedToken   RevokedToken `gorm:"not null;foreignKey:RevokedTokenId;references:ID;onDelete:CASCADE"`
	InstructorId   uint
	RevokedTokenId uint
}

func (rrti *RevokedTokenInstructor) BeforeSave(tx *gorm.DB) error {
	if err := checkExistence(tx, &Instructor{}, rrti.InstructorId); err != nil {
		return errors.New("instructor not found")
	}

	if err := checkExistence(tx, &RevokedToken{}, rrti.RevokedTokenId); err != nil {
		return errors.New("revoked token not found")
	}
	return nil
}

func (rrti *RevokedTokenInstructor) CreateRevokedTokenInstructor() (*RevokedTokenInstructor, error) {
	if err := db.Create(&rrti).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return rrti, nil
}

func GetAllRevokedTokenInstructor() ([]RevokedTokenInstructor, error) {
	var revokedTokenInstructors []RevokedTokenInstructor
	if err := db.Find(&revokedTokenInstructors).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return revokedTokenInstructors, nil
}

func GetRevokedTokenInstructorById(id uint) (*RevokedTokenInstructor, error) {
	var revokedTokenInstructor RevokedTokenInstructor
	if err := db.First(&revokedTokenInstructor, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedTokenInstructor, nil
}

func GetRevokedTokenInstructorByInstructorId(id uint) ([]RevokedTokenInstructor, error) {
	var revokedTokenInstructors []RevokedTokenInstructor
	if err := db.Where("instructor_id = ?", id).Find(&revokedTokenInstructors).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return revokedTokenInstructors, nil
}

func GetRevokedTokenInstructor(id uint, token string) (*RevokedTokenInstructor, error) {
	var revokedTokenInstructor RevokedTokenInstructor
	if err := db.Joins("JOIN revoked_tokens ON revoked_token_instructors.revoked_token_id = revoked_tokens.id").
		Where("revoked_token_instructors.instructor_id = ? AND revoked_tokens.token = ?", id, token).
		First(&revokedTokenInstructor).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedTokenInstructor, nil
}

func DeleteRevokedTokenInstructorById(id uint) (*RevokedTokenInstructor, error) {
	subQuery := db.Table("revoked_token_instructors").
		Where("revoked_token_instructors.id = ?", id).
		Select("revoked_token_instructors.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedTokenInstructor{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedTokenInstructor{}, nil
}

func DeleteRevokedTokenInstructorByInstructorId(id uint) ([]RevokedTokenInstructor, error) {
	subQuery := db.Table("revoked_token_instructors").
		Where("revoked_token_instructors.instructor_id = ?", id).
		Select("revoked_token_instructors.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedTokenInstructor{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return []RevokedTokenInstructor{}, nil
}

func DeleteRevokedTokenInstructor(id uint, token string) (*RevokedTokenInstructor, error) {
	subQuery := db.Table("revoked_token_instructors").
		Joins("JOIN revoked_tokens ON revoked_token_instructors.revoked_token_id = revoked_tokens.id").
		Where("revoked_token_instructors.instructor_id = ? AND revoked_tokens.refresh_token = ?", id, token).
		Select("revoked_token_instructors.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedTokenInstructor{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedTokenInstructor{}, nil
}
