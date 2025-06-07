package models

import (
	"gorm.io/gorm"
)

type RevokedRefreshToken struct {
	gorm.Model
	RefreshToken string `gorm:"not null"`
}

func (t *RevokedRefreshToken) CreateRevokedRefreshToken() (*RevokedRefreshToken, error) {
	if err := db.Create(&t).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return t, nil
}

func GetRevokedRefreshTokenById(id uint) (*RevokedRefreshToken, error) {
	var revokedRefreshToken RevokedRefreshToken
	if err := db.First(&revokedRefreshToken, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedRefreshToken, nil
}

func GetRevokedRefreshTokenByToken(token string) (*RevokedRefreshToken, error) {
	var revokedRefreshToken RevokedRefreshToken
	if err := db.Where("refresh_token = ?", token).First(&revokedRefreshToken).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedRefreshToken, nil
}

func DeleteRevokedRefreshTokenById(id uint) (*RevokedRefreshToken, error) {
	subQuery := db.Table("revoked_refresh_tokens").
		Where("revoked_refresh_tokens.id = ?", id).
		Select("revoked_refresh_tokens.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedRefreshToken{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedRefreshToken{}, nil
}
