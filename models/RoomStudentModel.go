package models

import (
	"errors"

	"gorm.io/gorm"
)

type RoomStudent struct {
	gorm.Model
	Student  Student `gorm:"not null;foreignKey:MemberId;onDelete:CASCADE"`
	Room     Room    `gorm:"not null;foreignKey:RoomId;onDelete:CASCADE"`
	MemberId uint
	RoomId   uint
}

func (rs *RoomStudent) BeforeSave(tx *gorm.DB) error {
	if err := checkExistence(tx, &Student{}, rs.MemberId); err != nil {
		return errors.New("student not found")
	}

	if err := checkExistence(tx, &Room{}, rs.RoomId); err != nil {
		return errors.New("room not found")
	}
	return nil
}

func (rs *RoomStudent) CreateRoomStudent() (*RoomStudent, error) {
	if err := db.Create(&rs).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return rs, nil
}

func GetAllRoomStudents() ([]RoomStudent, error) {
	var roomStudents []RoomStudent

	if err := db.Find(&roomStudents).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return roomStudents, nil
}
