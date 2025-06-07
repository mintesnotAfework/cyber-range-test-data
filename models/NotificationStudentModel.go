package models

import (
	"errors"

	"gorm.io/gorm"
)

type NotificationStudent struct {
	gorm.Model
	Student        Student      `gorm:"not null;foreignKey:StudentId;references:ID;onDelete:CASCADE"`
	Notification   Notification `gorm:"not null;foreignKey:NotificationId;references:ID;onDelete:CASCADE"`
	StudentId      uint
	NotificationId uint
}

func (ns *NotificationStudent) BeforeSave(tx *gorm.DB) error {
	if err := checkExistence(tx, &Student{}, ns.StudentId); err != nil {
		return errors.New("student not found")
	}

	if err := checkExistence(tx, &Notification{}, ns.NotificationId); err != nil {
		return errors.New("notification not found")
	}
	return nil
}

func (ns *NotificationStudent) CreateNotificationStudent() (*NotificationStudent, error) {
	if err := db.Create(&ns).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return ns, nil
}

func GetAllNotificationStudent() ([]NotificationStudent, error) {
	var notificationStudents []NotificationStudent
	if err := db.Find(&notificationStudents).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notificationStudents, nil
}

func GetNotificationStudentById(id uint) (*NotificationStudent, error) {
	var notificationStudent NotificationStudent
	if err := db.First(&notificationStudent, id).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notificationStudent, nil
}

func GetNotificationStudentByStudentId(id uint) ([]NotificationStudent, error) {
	var notificationStudents []NotificationStudent
	if err := db.Joins("JOIN students ON students.id = notification_students.student_id").
		Where("students.id=?", id).Find(&notificationStudents).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notificationStudents, nil
}

func GetNotificationByStudentId(id uint) ([]Notification, error) {
	var notifications []Notification
	if err := db.Joins("JOIN notification_students ON notifications.id = notification_students.notification_id").
		Joins("JOIN students ON students.id = notification_students.student_id").
		Where("students.id=?", id).Find(&notifications).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notifications, nil
}

func GetUnreadNotificationByStudentId(id uint) ([]Notification, error) {
	var notifications []Notification
	if err := db.Joins("JOIN notification_students ON notifications.id = notification_students.notification_id").
		Joins("JOIN students ON students.id = notification_students.student_id").
		Where("students.id=? AND notifications.read=?", id, false).Find(&notifications).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return notifications, nil
}

func GetNotificationStudent(id uint, content string) (*NotificationStudent, error) {
	var notificationStudent NotificationStudent
	if err := db.Joins("JOIN students ON students.id = notification_students.student_id").
		Joins("JOIN notifications ON notification_students.notification_id = notifications.id").
		Where("students.id=?", id).
		Where("notifications.content=?", content).First(&notificationStudent).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notificationStudent, nil
}

func GetNotificationByStudentIdAndNotificationId(student_id uint, notification_id uint) (*Notification, error) {
	var notification Notification
	if err := db.Joins("JOIN notification_students ON notifications.id = notification_students.notification_id").
		Joins("JOIN students ON students.id = notification_students.student_id").
		Where("students.id=?", student_id).
		Where("notifications.id=?", notification_id).
		First(&notification).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notification, nil
}

func GetNotificationCountByByStudentId(student_id uint) (*int64, error) {
	var count int64
	if err := db.Joins("JOIN notification_students ON notifications.id = notification_students.notification_id").
		Joins("JOIN students ON students.id = notification_students.student_id").
		Where("students.id=?", student_id).
		Model(&Notification{}).
		Count(&count).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &count, nil
}

func GetUnreadenNotificationCountByByStudentId(student_id uint) (*int64, error) {
	var count int64
	if err := db.Joins("JOIN notification_students ON notifications.id = notification_students.notification_id").
		Joins("JOIN students ON students.id = notification_students.student_id").
		Where("students.id=?", student_id).
		Where("notifications.read=?", false).
		Model(&Notification{}).
		Count(&count).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &count, nil
}

func GetReadenNotificationCountByStudentId(student_id uint) (*int64, error) {
	var count int64
	if err := db.Joins("JOIN notification_students ON notifications.id = notification_students.notification_id").
		Joins("JOIN students ON students.id = notification_students.student_id").
		Where("students.id=?", student_id).
		Where("notifications.read=?", true).
		Model(&Notification{}).
		Count(&count).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &count, nil
}

func DeleteNotificationStudentById(id uint) (*NotificationStudent, error) {
	var notificationStudent NotificationStudent
	if err := db.Where("id=?", id).Delete(&notificationStudent).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return &notificationStudent, nil
}

func DeleteNotificationStudentByStudentId(id uint) ([]NotificationStudent, error) {
	subQuery := db.Table("notification_students").
		Joins("JOIN students ON students.id = notification_students.student_id").
		Where("students.id=?", id).
		Select("notification_students.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&NotificationStudent{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return []NotificationStudent{}, nil
}

func DeleteNotificationStudent(id uint, content string) (*NotificationStudent, error) {
	subQuery := db.Table("notification_students").
		Joins("JOIN students ON students.id = notification_students.student_id").
		Where("students.id=? AND notifications.content=?", id, content).
		Select("notification_students.id")

	if err := db.Where("id IN (?)", subQuery).Delete(&NotificationStudent{}).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return &NotificationStudent{}, nil
}
