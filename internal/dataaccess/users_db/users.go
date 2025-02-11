package users_db

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) error
	UpdateUser(user *User) error
	// GetUserById(id uint) (*User, error)
	DeleteUserById(id uint) error
	// ListActiveUsers(start uint, end uint) ([]*User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

type User struct {
	ID          uint       `sql:"AUTO_INCREMENT" gorm:"primaryKey"`
	UserName    string     `gorm:"unique; no null"`
	FirstName   string     `gorm:"not null"`
	LastName    string     `gorm:"not null"`
	Email       string     `gorm:"not null"`
	Dob         *time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	MobilePhone string
	OfficePhone string
	Active      bool
}

func (r UserRepository) CreateUser(user *User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r UserRepository) UpdateUser(user *User) {
	r.DB.Save(user)
}

func (r UserRepository) DeleteUserById(id uint) error {
	result := r.DB.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
