package repo

import (
	"context"

	"github.com/moh-fajri/learn-jwt/repo/mysql"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

// User defines the user in db
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (u *User) GetWithEmail(ctx context.Context, email string) error {
	result := mysql.DB.WithContext(ctx).Where("email = ?", email).First(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Create data user
func (u *User) Create(ctx context.Context) error {
	result := mysql.DB.WithContext(ctx).Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// HashPassword encrypts user password
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword checks user password
func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
