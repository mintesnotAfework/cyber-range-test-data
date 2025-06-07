package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Title        string     `gorm:"not null;unique"`
	Description  string     `gorm:"not null;"`
	RoomVerified bool       `gorm:"not null:default:false"`
	Locked       bool       `gorm:"not null:default:false"`
	Instructor   Instructor `gorm:"not null;foreignKey:InstructorId;references:ID;onDelete:CASCADE" validate:"-"`
	InstructorId uint
	CreatedAt    time.Time
}

type RoomCreate struct {
	Title       string `validate:"required,min=2,max=50"`
	Description string `validate:"required,min=10,max=200"`
}

type RoomUpdate struct {
	Title       string `validate:"omitempty,min=2,max=50"`
	Description string `validate:"omitempty,min=10,max=200"`
}

func (r *Room) BeforeSave(tx *gorm.DB) error {
	r.Title = strings.ToLower(r.Title)
	return nil
}

func (r *Room) BeforeCreate(tx *gorm.DB) error {
	var existingInstructor Instructor
	result := tx.First(&existingInstructor, r.InstructorId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("instructor not found")
		}
		return result.Error
	}

	rc := &RoomCreate{
		Title:       strings.ToLower(r.Title),
		Description: r.Description,
	}

	validate := validateInit()
	err := validate.Struct(rc)
	if err != nil {
		return err
	}
	return nil
}

func (r *Room) BeforeUpdate(tx *gorm.DB) error {
	ru := &RoomUpdate{
		Title:       strings.ToLower(r.Title),
		Description: r.Description,
	}

	validate := validateInit()
	err := validate.Struct(ru)
	if err != nil {
		return err
	}
	return nil
}

func (r *Room) CreateRoom() (*Room, error) {
	if err := db.Create(&r).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}
	return r, nil
}

func GetAllRooms() ([]Room, error) {
	var rooms []Room

	if err := db.Find(&rooms).Error; err != nil {
		return nil, CustomErrorHandlerMiddleware(err)
	}

	return rooms, nil
}
