package models

import (
	"errors"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	User   User `gorm:"not null;foreignKey:UserId;references:ID;onDelete:CASCADE"`
	UserId uint `gorm:"uniqueIndex:idx_admin_user;"`
}

func (a *Admin) BeforeSave(tx *gorm.DB) error {
	if err := checkExistence(tx, &User{}, a.UserId); err != nil {
		return errors.New("user not found")
	}

	return nil
}

func (a *Admin) CreateAdmin() (*Admin, error) {
	var count int64
	if err := db.Model(&Admin{}).Count(&count).Error; err != nil || count != 0 {
		return nil, errors.New("admin already exists")
	}

	if err := db.Create(&a).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return a, nil
}

func GetAllAdmin() ([]Admin, error) {
	var admins []Admin
	if err := db.Find(&admins).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return admins, nil
}

func GetAdminById(id uint) (*Admin, error) {
	var admin Admin
	if err := db.First(&admin, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &admin, nil
}

func GetAdminByUserId(id uint) (*Admin, error) {
	var admin Admin
	if err := db.Joins("JOIN users ON users.id = admins.user_id").Where("users.id=?", id).First(&admin).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &admin, nil
}

func GetAdminByUserEmail(email string) (*Admin, error) {
	var admin Admin
	if err := db.Joins("JOIN users ON users.id = admins.user_id").Where("users.email=?", email).First(&admin).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &admin, nil
}

func DeleteAdminById(id int64) (*Admin, error) {
	subQuery := db.Table("admins").
		Where("admins.id = ?", id).
		Select("admins.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&Admin{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &Admin{}, nil
}

func DeleteAdminByUserId(id int64) (*Admin, error) {
	subQuery := db.Table("admins").
		Joins("JOIN users ON users.id = admins.user_id").
		Where("users.id = ?", id).
		Select("admins.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&Admin{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &Admin{}, nil
}
