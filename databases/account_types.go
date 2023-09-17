package databases

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID                int        `json:"id"`
	UserName          string     `json:"userName"`
	Email             string     `json:"email"`
	FirstName         string     `json:"firstName"`
	LastName          string     `json:"lastName"`
	EncryptedPassword string     `json:"-"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         *time.Time `json:"updatedAt"` // passing a pointer to handle null time cases
	DeletedAt         *time.Time `json:"deletedAt"` // passing a pointer to handle null time cases
}

type CreateAccountRequest struct {
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type UpdateAccountRequest struct {
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type DeleteAccountResponse struct {
	Message string `json:"message"`
}

func NewAccount(
	userName string,
	email string,
	firstName string,
	lastName string,
	password string,
) (*Account, error) {
	encryptPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		UserName:          userName,
		Email:             email,
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encryptPw),
		CreatedAt:         time.Now().UTC(),
	}, nil
}

func UpdateAccount(
	userName string,
	email string,
	firstName string,
	lastName string,
	password string,
) (*Account, error) {
	encryptPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	time := time.Now().UTC()
	return &Account{
		UserName:          userName,
		Email:             email,
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encryptPw),
		UpdatedAt:         &time, // passing a pointer to handle variable
	}, nil
}

func DeleteAccount() (*Account, error) {
	time := time.Now().UTC()
	return &Account{
		DeletedAt: &time, // passing a pointer to handle variable
	}, nil
}
