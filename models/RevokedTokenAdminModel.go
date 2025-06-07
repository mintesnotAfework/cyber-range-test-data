package models

import (
	"errors"

	"gorm.io/gorm"
)

type RevokedTokenAdmin struct {
	gorm.Model
	Admin          Admin        `gorm:"not null;foreignKey:AdminId;references:ID;onDelete:CASCADE"`
	RevokedToken   RevokedToken `gorm:"not null;foreignKey:RevokedTokenId;references:ID;onDelete:CASCADE"`
	AdminId        uint
	RevokedTokenId uint
}

func (rta *RevokedTokenAdmin) BeforeSave(tx *gorm.DB) error {
	if err := checkExistence(tx, &Admin{}, rta.AdminId); err != nil {
		return errors.New("admin not found")
	}

	if err := checkExistence(tx, &RevokedToken{}, rta.RevokedTokenId); err != nil {
		return errors.New("revoked token not found")
	}
	return nil
}

func (rta *RevokedTokenAdmin) CreateRevokedTokenAdmin() (*RevokedTokenAdmin, error) {
	if err := db.Create(&rta).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return rta, nil
}

func GetAllRevokedTokenAdmin() ([]RevokedTokenAdmin, error) {
	var revokedTokenAdmins []RevokedTokenAdmin
	if err := db.Find(&revokedTokenAdmins).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return revokedTokenAdmins, nil
}

func GetRevokedTokenAdminById(id uint) (*RevokedTokenAdmin, error) {
	var revokedTokenAdmin RevokedTokenAdmin
	if err := db.First(&revokedTokenAdmin, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedTokenAdmin, nil
}

func GetRevokedTokenAdminByAdminId(id uint) ([]RevokedTokenAdmin, error) {
	var revokedTokenAdmins []RevokedTokenAdmin
	if err := db.Where("admin_id = ?", id).Find(&revokedTokenAdmins).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return revokedTokenAdmins, nil
}

func GetRevokedTokenAdmin(id uint, token string) (*RevokedTokenAdmin, error) {
	var revokedTokenAdmin RevokedTokenAdmin
	if err := db.Joins("JOIN revoked_tokens ON revoked_token_admins.revoked_token_id = revoked_tokens.id").
		Where("revoked_token_admins.admin_id = ? AND revoked_tokens.token = ?", id, token).
		First(&revokedTokenAdmin).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedTokenAdmin, nil
}

func DeleteRevokedTokenAdminById(id uint) (*RevokedTokenAdmin, error) {
	subQuery := db.Table("revoked_token_admins").
		Where("revoked_token_admins.id = ?", id).
		Select("revoked_token_admins.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedTokenAdmin{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedTokenAdmin{}, nil
}

func DeleteRevokedTokenAdminByAdminId(id uint) ([]RevokedTokenAdmin, error) {
	subQuery := db.Table("revoked_token_admins").
		Where("revoked_token_admins.admin_id = ?", id).
		Select("revoked_token_admins.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedTokenAdmin{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return []RevokedTokenAdmin{}, nil
}

func DeleteRevokedTokenAdmin(id uint, token string) (*RevokedTokenAdmin, error) {
	subQuery := db.Table("revoked_token_admins").
		Joins("JOIN revoked_tokens ON revoked_token_admins.revoked_token_id = revoked_tokens.id").
		Where("revoked_token_admins.admin_id = ? AND revoked_tokens.token = ?", id, token).
		Select("revoked_token_admins.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedTokenAdmin{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedTokenAdmin{}, nil
}
