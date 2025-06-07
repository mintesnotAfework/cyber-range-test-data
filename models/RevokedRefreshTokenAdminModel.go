package models

import (
	"errors"

	"gorm.io/gorm"
)

type RevokedRefreshTokenAdmin struct {
	gorm.Model
	Admin                 Admin               `gorm:"not null;foreignKey:AdminId;references:ID;onDelete:CASCADE"`
	RevokedRefreshToken   RevokedRefreshToken `gorm:"not null;foreignKey:RevokedRefreshTokenId;references:ID;onDelete:CASCADE"`
	AdminId               uint
	RevokedRefreshTokenId uint
}

func (rrta *RevokedRefreshTokenAdmin) BeforeSave(tx *gorm.DB) error {
	if err := checkExistence(tx, &Admin{}, rrta.AdminId); err != nil {
		return errors.New("admin not found")
	}

	if err := checkExistence(tx, &RevokedRefreshToken{}, rrta.RevokedRefreshTokenId); err != nil {
		return errors.New("revoked refresh token not found")
	}
	return nil
}

func (rrta *RevokedRefreshTokenAdmin) CreateRevokedRefreshTokenAdmin() (*RevokedRefreshTokenAdmin, error) {
	if err := db.Create(&rrta).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return rrta, nil
}

func GetAllRevokedRefreshTokenAdmin() ([]RevokedRefreshTokenAdmin, error) {
	var revokedRefreshTokenAdmins []RevokedRefreshTokenAdmin
	if err := db.Find(&revokedRefreshTokenAdmins).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return revokedRefreshTokenAdmins, nil
}

func GetRevokedRefreshTokenAdminById(id uint) (*RevokedRefreshTokenAdmin, error) {
	var revokedRefreshTokenAdmin RevokedRefreshTokenAdmin
	if err := db.First(&revokedRefreshTokenAdmin, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedRefreshTokenAdmin, nil
}

func GetRevokedRefreshTokenAdminByAdminId(id uint) ([]RevokedRefreshTokenAdmin, error) {
	var revokedRefreshTokenAdmins []RevokedRefreshTokenAdmin
	if err := db.Where("admin_id = ?", id).Find(&revokedRefreshTokenAdmins).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return revokedRefreshTokenAdmins, nil
}

func GetRevokedRefreshTokenAdmin(id uint, token string) (*RevokedRefreshTokenAdmin, error) {
	var revokedRefreshTokenAdmin RevokedRefreshTokenAdmin
	if err := db.Joins("JOIN revoked_refresh_tokens ON revoked_refresh_token_admins.revoked_refresh_token_id = revoked_refresh_tokens.id").
		Where("revoked_refresh_token_admins.admin_id = ? AND revoked_refresh_tokens.refresh_token = ?", id, token).
		First(&revokedRefreshTokenAdmin).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedRefreshTokenAdmin, nil
}

func DeleteRevokedRefreshTokenAdminById(id uint) (*RevokedRefreshTokenAdmin, error) {
	subQuery := db.Table("revoked_refresh_token_admins").
		Where("revoked_refresh_token_admins.id = ?", id).
		Select("revoked_refresh_token_admins.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedRefreshTokenAdmin{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedRefreshTokenAdmin{}, nil
}

func DeleteRevokedRefreshTokenAdminByAdminId(id uint) ([]RevokedRefreshTokenAdmin, error) {
	subQuery := db.Table("revoked_refresh_token_admins").
		Where("revoked_refresh_token_admins.admin_id = ?", id).
		Select("revoked_refresh_token_admins.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedRefreshTokenAdmin{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return []RevokedRefreshTokenAdmin{}, nil
}

func DeleteRevokedRefreshTokenAdmin(id uint, token string) (*RevokedRefreshTokenAdmin, error) {
	subQuery := db.Table("revoked_refresh_token_admins").
		Joins("JOIN revoked_refresh_tokens ON revoked_refresh_token_admins.revoked_refresh_token_id = revoked_refresh_tokens.id").
		Where("revoked_refresh_token_admins.admin_id = ? AND revoked_refresh_tokens.refresh_token = ?", id, token).
		Select("revoked_refresh_token_admins.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedRefreshTokenAdmin{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedRefreshTokenAdmin{}, nil
}
