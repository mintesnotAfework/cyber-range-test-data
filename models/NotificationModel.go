package models

import (
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	Message string    `gorm:"not null"`
	Type    string    `gorm:"not null"`
	SendAt  time.Time `gorm:"not null"`
	Read    bool      `gorm:"not null,default:false"`
}

type NotificationCreate struct {
	Message string `validate:"required"`
	Type    string `validate:"required,oneof='announcement' 'system' 'reminder'"`
}

type NotificationUpdate struct {
	Message string `validate:"omitempty"`
	Type    string `validate:"omitempty,oneof='announcement' 'system' 'reminder'"`
	Read    bool   `validate:"omitempty"`
}

func (n *Notification) BeforeSave(tx *gorm.DB) error {
	n.SendAt = time.Now()
	return nil
}
func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	n.Type = strings.ToLower(n.Type)

	nc := NotificationCreate{
		Message: n.Message,
		Type:    n.Type,
	}

	validate := validateInit()
	err := validate.Struct(nc)
	if err != nil {
		return err
	}

	return nil
}

func (n *Notification) BeforeUpdate(tx *gorm.DB) error {
	n.Type = strings.ToLower(n.Type)

	nu := NotificationUpdate{
		Message: n.Message,
		Type:    n.Type,
		Read:    n.Read}

	validate := validateInit()
	err := validate.Struct(nu)
	if err != nil {
		return err
	}

	return nil
}

func (n *Notification) CreateNotification() (*Notification, error) {
	if err := db.Create(&n).Error; err != nil {
		log.Println(err.Error())
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return n, nil
}

func GetAllNotification() ([]Notification, error) {
	var notifications []Notification
	if err := db.Find(&notifications).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notifications, nil
}

func GetNotificationById(id uint) (*Notification, error) {
	var notification Notification
	if err := db.First(&notification, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notification, nil
}

func (n *Notification) UpdateNotificationById(id uint) (*Notification, error) {
	var notification Notification
	if err := db.First(&notification, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	if err := db.Model(&notification).Update("message", n.Message).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &notification, nil
}

func (n *Notification) UpdateNotificationReadStatus(id uint) (*Notification, error) {
	var notification Notification
	if err := db.First(&notification, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	if err := db.Model(&notification).Update("read", true).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &notification, nil
}

func DeleteNotificationById(id uint) (*Notification, error) {
	subQuery := db.Table("notifications").
		Where("notifications.id = ?", id).
		Select("notifications.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&Notification{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &Notification{}, nil
}
