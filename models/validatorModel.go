package models

import (
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	hasUpper      = regexp.MustCompile(`[A-Z]`)
	hasLower      = regexp.MustCompile(`[a-z]`)
	hasNumber     = regexp.MustCompile(`[0-9]`)
	hasSpecial    = regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`)
	phoneRegex    = regexp.MustCompile(`^\+[0-9]\d{1,14}$`) // E.164 international format
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
)

func validateInit() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("password", validatePassword)
	validate.RegisterValidation("phone", validatePhone)
	validate.RegisterValidation("username", validateUsername)
	return validate
}

func validateUsername(f1 validator.FieldLevel) bool {
	return usernameRegex.MatchString(f1.Field().String())
}
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Check if all rules are satisfied
	return hasUpper.MatchString(password) &&
		hasLower.MatchString(password) &&
		hasNumber.MatchString(password) &&
		hasSpecial.MatchString(password) &&
		len(password) >= 8
}

func validatePhone(fl validator.FieldLevel) bool {
	return phoneRegex.MatchString(fl.Field().String())
}

func checkExistence(tx *gorm.DB, model interface{}, id uint) error {
	if err := tx.First(model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return nil
}
