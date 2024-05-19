package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID `gorm:"type:uuid"`
	Name            string
	EmailAddress    string `gorm:"unique"`
	Password        string
	Locked          bool
	LoginAttempts   uint
	UserPermissions []UserPermission
}

type UserPermission struct {
	ID           uuid.UUID `gorm:"type:uuid"`
	UserID       uuid.UUID `gorm:"type:uuid"`
	PermissionID uuid.UUID `gorm:"type:uuid"`
	Permission   Permission
	RestaurantID *uuid.UUID `gorm:"type:uuid"`
}

type UserData struct {
	Db *gorm.DB
}

func NewUserData(Db *gorm.DB) *UserData {
	return &UserData{Db: Db}
}

// Create user.
func (userData *UserData) CreateUser(user User) (*User, error) {
	result := userData.Db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Update user.
func (userData *UserData) UpdateUser(user User) (*User, error) {
	result := userData.Db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Get user by email address.
func (userData *UserData) GetUserByEmailAddress(emailAddress string, user User) (*User, error) {
	result := userData.Db.Where("email_address = ?", emailAddress).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
