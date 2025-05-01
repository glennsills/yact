package userData

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) error
	UpdateUser(user *User) error
	// GetUserById(id uint) (*User, error)
	DeleteUserById(id uint) error
	ListActiveUsers(start uint, end uint) ([]*User, error)
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

func (r UserRepository) CreateUser(email string, firstName string, lastName string, dateOfBirth string, phone string, officePhone string) (*User, error) {

	// TODO: Handle field validation better

	if len(dateOfBirth) == 10 && len(phone) == 10 && (len(officePhone) == 10 || len(officePhone) == 0) {
		dob, err := time.Parse("2006-01-02", dateOfBirth)
		if err != nil {
			return nil, err
		} else {
			user := User{
				UserName:    email,
				FirstName:   firstName,
				LastName:    lastName,
				Email:       email,
				Dob:         &dob,
				MobilePhone: phone,
				OfficePhone: officePhone,
			}
			result := r.DB.Create(user)
			if result.Error != nil {
				return nil, result.Error
			}
			return &user, nil
		}
	} else {
		return nil, errors.New("invalid input parameter")
	}
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

func ListActiveUsers(start uint, end uint) ([]*User, error) {
	return nil, nil
}
