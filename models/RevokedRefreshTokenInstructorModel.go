package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type RevokedRefreshTokenInstructor struct {
	gorm.Model
	Instructor            Instructor          `gorm:"not null;foreignKey:InstructorId;references:ID;onDelete:CASCADE"`
	RevokedRefreshToken   RevokedRefreshToken `gorm:"not null;foreignKey:RevokedRefreshTokenId;references:ID;onDelete:CASCADE"`
	InstructorId          uint
	RevokedRefreshTokenId uint
}

func (rrti *RevokedRefreshTokenInstructor) BeforeSave(tx *gorm.DB) error {
	if err := checkExistence(tx, &Instructor{}, rrti.InstructorId); err != nil {
		return errors.New("instructor not found")
	}

	if err := checkExistence(tx, &RevokedRefreshToken{}, rrti.RevokedRefreshTokenId); err != nil {
		return errors.New("revoked refresh token not found")
	}
	return nil
}

func (rrti *RevokedRefreshTokenInstructor) CreateRevokedRefreshTokenInstructor() (*RevokedRefreshTokenInstructor, error) {
	if err := db.Create(&rrti).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return rrti, nil
}

func GetAllRevokedRefreshTokenInstructor() ([]RevokedRefreshTokenInstructor, error) {
	var revokedRefreshTokenInstructors []RevokedRefreshTokenInstructor
	if err := db.Find(&revokedRefreshTokenInstructors).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return revokedRefreshTokenInstructors, nil
}

func GetRevokedRefreshTokenInstructorById(id uint) (*RevokedRefreshTokenInstructor, error) {
	var revokedRefreshTokenInstructor RevokedRefreshTokenInstructor
	if err := db.First(&revokedRefreshTokenInstructor, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedRefreshTokenInstructor, nil
}

func GetRevokedRefreshTokenInstructorByInstructorId(id uint) ([]RevokedRefreshTokenInstructor, error) {
	var revokedRefreshTokenInstructors []RevokedRefreshTokenInstructor
	if err := db.Where("instructor_id = ?", id).Find(&revokedRefreshTokenInstructors).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return revokedRefreshTokenInstructors, nil
}

func GetRevokedRefreshTokenInstructor(id uint, token string) (*RevokedRefreshTokenInstructor, error) {
	var revokedRefreshTokenInstructor RevokedRefreshTokenInstructor
	if err := db.Joins("JOIN revoked_refresh_tokens ON revoked_refresh_token_instructors.revoked_refresh_token_id = revoked_refresh_tokens.id").
		Where("revoked_refresh_token_instructors.instructor_id = ? AND revoked_refresh_tokens.refresh_token = ?", id, token).
		First(&revokedRefreshTokenInstructor).Error; err != nil {
		log.Println(err.Error())
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedRefreshTokenInstructor, nil
}

func DeleteRevokedRefreshTokenInstructorById(id uint) (*RevokedRefreshTokenInstructor, error) {
	subQuery := db.Table("revoked_refresh_token_instructors").
		Where("revoked_refresh_token_instructors.id = ?", id).
		Select("revoked_refresh_token_instructors.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedRefreshTokenInstructor{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedRefreshTokenInstructor{}, nil
}

func DeleteRevokedRefreshTokenInstructorByInstructorId(id uint) ([]RevokedRefreshTokenInstructor, error) {
	subQuery := db.Table("revoked_refresh_token_instructors").
		Where("revoked_refresh_token_instructors.instructor_id = ?", id).
		Select("revoked_refresh_token_instructors.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedRefreshTokenInstructor{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return []RevokedRefreshTokenInstructor{}, nil
}

func DeleteRevokedRefreshTokenInstructor(id uint, token string) (*RevokedRefreshTokenInstructor, error) {
	subQuery := db.Table("revoked_refresh_token_instructors").
		Joins("JOIN revoked_refresh_tokens ON revoked_refresh_token_instructors.revoked_refresh_token_id = revoked_refresh_tokens.id").
		Where("revoked_refresh_token_instructors.instructor_id = ? AND revoked_refresh_tokens.refresh_token = ?", id, token).
		Select("revoked_refresh_token_instructors.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedRefreshTokenInstructor{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedRefreshTokenInstructor{}, nil
}
