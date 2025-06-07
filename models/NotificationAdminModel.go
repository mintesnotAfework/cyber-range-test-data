package models

import (
	"errors"

	"gorm.io/gorm"
)

type NotificationAdmin struct {
	gorm.Model
	Admin          Admin        `gorm:"not null;foreignKey:AdminId;references:ID;onDelete:CASCADE"`
	Notification   Notification `gorm:"not null;foreignKey:NotificationId;references:ID;onDelete:CASCADE"`
	AdminId        uint
	NotificationId uint
}

func (na *NotificationAdmin) BeforeCreate(tx *gorm.DB) error {
	if err := checkExistence(tx, &Admin{}, na.AdminId); err != nil {
		return errors.New("admin not found")
	}

	if err := checkExistence(tx, &Notification{}, na.NotificationId); err != nil {
		return errors.New("notification not found")
	}
	return nil
}

func (na *NotificationAdmin) CreateNotificationAdmin() (*NotificationAdmin, error) {
	if err := db.Create(&na).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return na, nil
}

func GetAllNotificationAdmin() ([]NotificationAdmin, error) {
	var notificationAdmins []NotificationAdmin
	if err := db.Find(&notificationAdmins).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notificationAdmins, nil
}

func GetNotificationAdminById(id uint) (*NotificationAdmin, error) {
	var notificationAdmin NotificationAdmin
	if err := db.First(&notificationAdmin, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notificationAdmin, nil
}

func GetNotificationAdminByAdminId(id uint) ([]NotificationAdmin, error) {
	var notificationAdmins []NotificationAdmin
	if err := db.Joins("JOIN admins ON admins.id = notification_admins.admin_id").
		Where("admins.id=?", id).Find(&notificationAdmins).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notificationAdmins, nil
}

func GetNotificationByAdminId(id uint) ([]Notification, error) {
	var notifications []Notification
	if err := db.Joins("JOIN notification_admins ON notifications.id = notification_admins.notification_id").
		Joins("JOIN admins ON admins.id = notification_admins.admin_id").
		Where("admins.id=?", id).Find(&notifications).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notifications, nil
}

func GetUnreadNotificationByAdminId(id uint) ([]Notification, error) {
	var notifications []Notification
	if err := db.Joins("JOIN notification_admins ON notifications.id = notification_admins.notification_id").
		Joins("JOIN admins ON admins.id = notification_admins.admin_id").
		Where("admins.id=? AND notifications.read=?", id, false).Find(&notifications).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notifications, nil
}

func GetNotificationByAdminIdAndNotificationId(admin_id uint, notification_id uint) (*Notification, error) {
	var notification Notification
	if err := db.Joins("JOIN notification_admins ON notifications.id = notification_admins.notification_id").
		Joins("JOIN admins ON admins.id = notification_admins.admin_id").
		Where("admins.id=?", admin_id).
		Where("notifications.id=?", notification_id).First(&notification).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notification, nil
}

func GetNotificationCountByAdminId(admin_id uint) (*int64, error) {
	var count int64
	if err := db.Joins("JOIN notification_admins ON notifications.id = notification_admins.notification_id").
		Joins("JOIN admins ON admins.id = notification_admins.admin_id").
		Where("admins.id=?", admin_id).
		Model(&Notification{}).
		Count(&count).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &count, nil
}

func GetUnreadenNotificationCountByAdminId(admin_id uint) (*int64, error) {
	var count int64
	if err := db.Joins("JOIN notification_admins ON notifications.id = notification_admins.notification_id").
		Joins("JOIN admins ON admins.id = notification_admins.admin_id").
		Where("admins.id=?", admin_id).
		Where("notifications.read=?", false).
		Model(&Notification{}).
		Count(&count).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &count, nil
}

func GetReadenNotificationCountByAdminId(admin_id uint) (*int64, error) {
	var count int64
	if err := db.Joins("JOIN notification_admins ON notifications.id = notification_admins.notification_id").
		Joins("JOIN admins ON admins.id = notification_admins.admin_id").
		Where("admins.id=?", admin_id).
		Where("notifications.read=?", true).
		Model(&Notification{}).
		Count(&count).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &count, nil
}

func GetNotificationAdmin(id uint, content string) (*NotificationAdmin, error) {
	var notificationAdmin NotificationAdmin
	if err := db.Joins("JOIN admins ON admins.id = notification_admins.admin_id").
		Joins("JOIN notifications ON notification_admins.notification_id = notifications.id").
		Where("admins.id=?", id).
		Where("notifications.content=?", content).First(&notificationAdmin).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notificationAdmin, nil
}

func DeleteNotificationAdminById(id uint) (*NotificationAdmin, error) {
	var notificationAdmin NotificationAdmin
	if err := db.Where("id=?", id).Delete(&notificationAdmin).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notificationAdmin, nil
}

func DeleteNotificationAdminByAdminId(id uint) ([]NotificationAdmin, error) {
	subQuery := db.Table("notification_admins").
		Joins("JOIN admins ON admins.id = notification_admins.admin_id").
		Where("admins.id=?", id).
		Select("notification_admins.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&NotificationAdmin{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return []NotificationAdmin{}, nil

}

func DeleteNotificationAdmin(id uint, content string) (*NotificationAdmin, error) {
	subQuery := db.Table("notification_admins").
		Joins("JOIN admins ON admins.id = notification_admins.admin_id").
		Where("admins.id=? AND notifications.content=?", id, content).
		Select("notification_admins.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&NotificationAdmin{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &NotificationAdmin{}, nil
}
