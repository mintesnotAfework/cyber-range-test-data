package models

import (
	"errors"

	"gorm.io/gorm"
)

type RevokedRefreshTokenStudent struct {
	gorm.Model
	Student               Student             `gorm:"not null;foreignKey:StudentId;references:ID;onDelete:CASCADE"`
	RevokedRefreshToken   RevokedRefreshToken `gorm:"not null;foreignKey:RevokedRefreshTokenId;references:ID;onDelete:CASCADE"`
	StudentId             uint
	RevokedRefreshTokenId uint
}

func (rrts *RevokedRefreshTokenStudent) BeforeSave(tx *gorm.DB) error {
	if err := checkExistence(tx, &Student{}, rrts.StudentId); err != nil {
		return errors.New("student not found")
	}

	if err := checkExistence(tx, &RevokedRefreshToken{}, rrts.RevokedRefreshTokenId); err != nil {
		return errors.New("revoked refresh token not found")
	}
	return nil
}

func (rrts *RevokedRefreshTokenStudent) CreateRevokedRefreshTokenStudent() (*RevokedRefreshTokenStudent, error) {
	if err := db.Create(&rrts).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return rrts, nil
}

func GetAllRevokedRefreshTokenStudent() ([]RevokedRefreshTokenStudent, error) {
	var revokedRefreshTokenStudents []RevokedRefreshTokenStudent
	if err := db.Find(&revokedRefreshTokenStudents).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return revokedRefreshTokenStudents, nil
}

func GetRevokedRefreshTokenStudentById(id uint) (*RevokedRefreshTokenStudent, error) {
	var revokedRefreshTokenStudent RevokedRefreshTokenStudent
	if err := db.First(&revokedRefreshTokenStudent, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedRefreshTokenStudent, nil
}

func GetRevokedRefreshTokenStudentByStudentId(id uint) ([]RevokedRefreshTokenStudent, error) {
	var revokedRefreshTokenStudents []RevokedRefreshTokenStudent
	if err := db.Where("student_id = ?", id).Find(&revokedRefreshTokenStudents).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return revokedRefreshTokenStudents, nil
}

func GetRevokedRefreshTokenStudent(id uint, token string) (*RevokedRefreshTokenStudent, error) {
	var revokedRefreshTokenStudent RevokedRefreshTokenStudent
	if err := db.Joins("JOIN revoked_refresh_tokens ON revoked_refresh_token_students.revoked_refresh_token_id = revoked_refresh_tokens.id").
		Where("revoked_refresh_token_students.student_id = ? AND revoked_refresh_tokens.refresh_token = ?", id, token).
		First(&revokedRefreshTokenStudent).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedRefreshTokenStudent, nil
}

func DeleteRevokedRefreshTokenStudentById(id uint) (*RevokedRefreshTokenStudent, error) {
	subQuery := db.Table("revoked_refresh_token_students").
		Where("revoked_refresh_token_students.id = ?", id).
		Select("revoked_refresh_token_students.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedRefreshTokenStudent{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedRefreshTokenStudent{}, nil
}

func DeleteRevokedRefreshTokenStudentByStudentId(id uint) ([]RevokedRefreshTokenStudent, error) {
	subQuery := db.Table("revoked_refresh_token_students").
		Where("revoked_refresh_token_students.student_id = ?", id).
		Select("revoked_refresh_token_students.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedRefreshTokenStudent{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return []RevokedRefreshTokenStudent{}, nil
}

func DeleteRevokedRefreshTokenStudent(id uint, token string) (*RevokedRefreshTokenStudent, error) {
	subQuery := db.Table("revoked_refresh_token_students").
		Joins("JOIN revoked_refresh_tokens ON revoked_refresh_token_students.revoked_refresh_token_id = revoked_refresh_tokens.id").
		Where("revoked_refresh_token_students.student_id = ? AND revoked_refresh_tokens.refresh_token = ?", id, token).
		Select("revoked_refresh_token_students.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedRefreshTokenStudent{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedRefreshTokenStudent{}, nil
}
