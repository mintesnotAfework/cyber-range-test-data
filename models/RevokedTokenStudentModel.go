package models

import (
	"errors"

	"gorm.io/gorm"
)

type RevokedTokenStudent struct {
	gorm.Model
	Student        Student      `gorm:"not null;foreignKey:StudentId;references:ID;onDelete:CASCADE"`
	RevokedToken   RevokedToken `gorm:"not null;foreignKey:RevokedTokenId;references:ID;onDelete:CASCADE"`
	StudentId      uint
	RevokedTokenId uint
}

func (rts *RevokedTokenStudent) BeforeSave(tx *gorm.DB) error {
	if err := checkExistence(tx, &Student{}, rts.StudentId); err != nil {
		return errors.New("student not found")
	}

	if err := checkExistence(tx, &RevokedToken{}, rts.RevokedTokenId); err != nil {
		return errors.New("revoked token not found")
	}
	return nil
}

func (rts *RevokedTokenStudent) CreateRevokedTokenStudent() (*RevokedTokenStudent, error) {
	if err := db.Create(&rts).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return rts, nil
}

func GetAllRevokedTokenStudent() ([]RevokedTokenStudent, error) {
	var revokedTokenStudents []RevokedTokenStudent
	if err := db.Find(&revokedTokenStudents).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return revokedTokenStudents, nil
}

func GetRevokedTokenStudentById(id uint) (*RevokedTokenStudent, error) {
	var revokedTokenStudent RevokedTokenStudent
	if err := db.First(&revokedTokenStudent, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedTokenStudent, nil
}

func GetRevokedTokenStudentByStudentId(id uint) ([]RevokedTokenStudent, error) {
	var revokedTokenStudents []RevokedTokenStudent
	if err := db.Where("student_id = ?", id).Find(&revokedTokenStudents).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return revokedTokenStudents, nil
}

func GetRevokedTokenStudent(id uint, token string) (*RevokedTokenStudent, error) {
	var revokedTokenStudent RevokedTokenStudent
	if err := db.Joins("JOIN revoked_tokens ON revoked_token_students.revoked_token_id = revoked_tokens.id").
		Where("revoked_token_students.student_id = ? AND revoked_tokens.token = ?", id, token).
		First(&revokedTokenStudent).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedTokenStudent, nil
}

func DeleteRevokedTokenStudentById(id uint) (*RevokedTokenStudent, error) {
	subQuery := db.Table("revoked_token_students").
		Where("revoked_token_students.id = ?", id).
		Select("revoked_token_students.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedTokenStudent{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedTokenStudent{}, nil
}

func DeleteRevokedTokenStudentByStudentId(id uint) ([]RevokedTokenStudent, error) {
	subQuery := db.Table("revoked_token_students").
		Where("revoked_token_students.student_id = ?", id).
		Select("revoked_token_students.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedTokenStudent{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return []RevokedTokenStudent{}, nil
}

func DeleteRevokedTokenStudent(id uint, token string) (*RevokedTokenStudent, error) {
	subQuery := db.Table("revoked_token_students").
		Joins("JOIN revoked_tokens ON revoked_token_students.revoked_token_id = revoked_tokens.id").
		Where("revoked_token_students.student_id = ? AND revoked_tokens.token = ?", id, token).
		Select("revoked_token_students.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedTokenStudent{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedTokenStudent{}, nil
}
