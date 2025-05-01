package userService

import (
	"time"

	usersData "github.com/glennsills/yact/db/usersData"
)

type Service interface {
	ListActiveUsers() (*UserListing, error)
	FindUser(emailFilter string, onlyActive bool) ([]*UserListing, error)
}

type User struct {
	ID          int
	Email       string
	FirstName   string
	LastName    string
	Phone       string
	OfficePhone string
	Dob         string
	UpdatedAt   time.Time
}

type UserListItem struct {
	ID          int
	Email       string
	FirstName   string
	LastName    string
	Phone       string
	OfficePhone string
}

type UserListing struct {
	List     *[]UserListItem
	NextPage int
	PrevPage int
}

func NewUserService(dataRepo *usersData.UserRepository) *UserService {
	return &UserService{
		DataRepo: *dataRepo,
	}
}

type UserService struct {
	DataRepo usersData.UserRepository
}

func (service UserService) ListActiveUsers() (*UserListing, error) {
	users := &UserListing{
		List: &[]UserListItem{
			{
				ID:          1,
				Email:       "john.doe@example.com",
				FirstName:   "John",
				LastName:    "Doe",
				Phone:       "123-456-7890",
				OfficePhone: "098-765-4321",
			},
			{
				ID:          2,
				Email:       "jane.smith@example.com",
				FirstName:   "Jane",
				LastName:    "Smith",
				Phone:       "234-567-8901",
				OfficePhone: "987-654-3210",
			},
		},
		NextPage: 2,
		PrevPage: 1,
	}
	return users, nil
}

func (service UserService) FindUser(emailFilter string, onlyActive bool) ([]*UserListing, error) {
	return nil, nil
}

func (service UserService) NewUser(email string, firstName string, lastName string, phone string, officePhone string, dob string) (*User, error) {

	user, err := service.DataRepo.CreateUser(email, firstName, lastName, dob, phone, officePhone)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        int(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Dob:       user.Dob.Format("2006-01-02"),
		UpdatedAt: user.CreatedAt,
	}, nil
}
