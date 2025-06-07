package models

import (
	"errors"

	"gorm.io/gorm"
)

type NotificationInstructor struct {
	gorm.Model
	Instructor     Instructor   `gorm:"not null;foreignKey:InstructorId;references:ID;onDelete:CASCADE"`
	Notification   Notification `gorm:"not null;foreignKey:NotificationId;references:ID;onDelete:CASCADE"`
	InstructorId   uint
	NotificationId uint
}

func (ni *NotificationInstructor) BeforeCreate(tx *gorm.DB) error {
	if err := checkExistence(tx, &Instructor{}, ni.InstructorId); err != nil {
		return errors.New("instructor not found")
	}

	if err := checkExistence(tx, &Notification{}, ni.NotificationId); err != nil {
		return errors.New("notification not found")
	}
	return nil
}

func (ni *NotificationInstructor) CreateNotificationInstructor() (*NotificationInstructor, error) {
	if err := db.Create(&ni).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return ni, nil
}

func GetAllNotificationInstructor() ([]NotificationInstructor, error) {
	var notificationInstructors []NotificationInstructor
	if err := db.Find(&notificationInstructors).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notificationInstructors, nil
}

func GetNotificationInstructorById(id uint) (*NotificationInstructor, error) {
	var notificationInstructor NotificationInstructor
	if err := db.First(&notificationInstructor, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notificationInstructor, nil
}

func GetNotificationInstructorByInstructorId(id uint) ([]NotificationInstructor, error) {
	var notificationInstructors []NotificationInstructor
	if err := db.Joins("JOIN instructors ON instructors.id = notification_instructors.instructor_id").
		Where("instructors.id=?", id).Find(&notificationInstructors).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notificationInstructors, nil
}

func GetNotificationByInstructorId(id uint) ([]Notification, error) {
	var notifications []Notification
	if err := db.Joins("JOIN notification_instructors ON notifications.id = notification_instructors.notification_id").
		Joins("JOIN instructors ON instructors.id = notification_instructors.instructor_id").
		Where("instructors.id=?", id).Find(&notifications).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notifications, nil
}

func GetUnreadNotificationByInstructorId(id uint) ([]Notification, error) {
	var notifications []Notification
	if err := db.Joins("JOIN notification_instructors ON notifications.id = notification_instructors.notification_id").
		Joins("JOIN instructors ON instructors.id = notification_instructors.instructor_id").
		Where("instructors.id=? AND notifications.read=?", id, false).Find(&notifications).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notifications, nil
}

func GetNotificationInstructor(id uint, content string) (*NotificationInstructor, error) {
	var notificationInstructor NotificationInstructor
	if err := db.Joins("JOIN instructors ON instructors.id = notification_instructors.instructor_id").
		Joins("JOIN notifications ON notification_instructors.notification_id = notifications.id").
		Where("instructors.id=?", id).
		Where("notifications.content=?", content).First(&notificationInstructor).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notificationInstructor, nil
}

func DeleteNotificationInstructorById(id uint) (*NotificationInstructor, error) {
	var notificationInstructor NotificationInstructor
	if err := db.Where("id=?", id).Delete(&notificationInstructor).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notificationInstructor, nil
}

func GetNotificationByInstructorIdAndNotificationId(instructor_id uint, notification_id uint) (*Notification, error) {
	var notification Notification
	if err := db.Joins("JOIN notification_instructors ON notifications.id = notification_instructors.notification_id").
		Joins("JOIN instructors ON instructors.id = notification_instructors.instructor_id").
		Where("instructors.id=?", instructor_id).
		Where("notifications.id=?", notification_id).
		First(&notification).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notification, nil
}

func GetNotificationCountByInstructorId(instructor_id uint) (*int64, error) {
	var count int64
	if err := db.Joins("JOIN notification_instructors ON notifications.id = notification_instructors.notification_id").
		Joins("JOIN instructors ON instructors.id = notification_instructors.instructor_id").
		Where("instructors.id=?", instructor_id).
		Model(&Notification{}).
		Count(&count).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &count, nil
}

func GetUnreadenNotificationCountByInstructorId(instructor_id uint) (*int64, error) {
	var count int64
	if err := db.Joins("JOIN notification_instructors ON notifications.id = notification_instructors.notification_id").
		Joins("JOIN instructors ON instructors.id = notification_instructors.instructor_id").
		Where("instructors.id=?", instructor_id).
		Where("notifications.read=?", false).
		Model(&Notification{}).
		Count(&count).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &count, nil
}

func GetReadenNotificationCountByInstructorId(instructor_id uint) (*int64, error) {
	var count int64
	if err := db.Joins("JOIN notification_instructors ON notifications.id = notification_instructors.notification_id").
		Joins("JOIN instructors ON instructors.id = notification_instructors.instructor_id").
		Where("instructors.id=?", instructor_id).
		Where("notifications.read=?", true).
		Model(&Notification{}).
		Count(&count).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &count, nil
}

func DeleteNotificationInstructorByInstructorId(id uint) ([]NotificationInstructor, error) {
	subQuery := db.Table("notification_instructors").
		Joins("JOIN instructors ON instructors.id = notification_instructors.instructor_id").
		Where("instructors.id=?", id).
		Select("notification_instructors.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&NotificationInstructor{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return []NotificationInstructor{}, nil
}

func DeleteNotificationInstructor(id uint, content string) (*NotificationInstructor, error) {
	subQuery := db.Table("notification_instructors").
		Joins("JOIN instructors ON instructors.id = notification_instructors.instructor_id").
		Where("instructors.id=? AND notifications.content=?", id, content).
		Select("notification_instructors.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&NotificationInstructor{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &NotificationInstructor{}, nil
}
