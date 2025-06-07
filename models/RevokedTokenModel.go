package models

import (
	"gorm.io/gorm"
)

type RevokedToken struct {
	gorm.Model
	Token string `gorm:"not null"`
}

func (t *RevokedToken) CreateRevokedToken() (*RevokedToken, error) {
	if err := db.Create(&t).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return t, nil
}

func GetRevokedTokenById(id uint) (*RevokedToken, error) {
	var revokedToken RevokedToken
	if err := db.First(&revokedToken, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedToken, nil
}

func GetRevokedTokenByToken(token string) (*RevokedToken, error) {
	var revokedToken RevokedToken
	if err := db.Where("token = ?", token).First(&revokedToken).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &revokedToken, nil
}

func DeleteRevokedTokenById(id uint) (*RevokedToken, error) {
	subQuery := db.Table("revoked_tokens").
		Where("revoked_tokens.id = ?", id).
		Select("revoked_tokens.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&RevokedToken{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &RevokedToken{}, nil
}
