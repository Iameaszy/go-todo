package AuthRepository

import (
	"errors"
	"todo/models"
	"todo/models/db"

	"gorm.io/gorm"
)

func GetUser(user *models.User, email string) (err error) {
	err = db.DB.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}

func CreateUser(user *models.User) (err error) {
	if err = db.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
