package entity

import (
	"errors"
	"math/rand"
	"time"
)

// letterRunes is a list of runes that can be used to generate a random string.
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// User is a struct that represents a user.
type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	CreatedDate string `json:"createdDate"`
	UpdatedDate string `json:"updatedDate"`
}

// UserCreation is a struct that is used to create a new user.
type UserCreation struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

/*
MapToUser maps the UserCreation struct to a User struct.
If any of the required fields are missing, it will return an error.
Otherwise, it will return the created user.
*/
func (user UserCreation) MapToUser() (User, error) {
	if user.Password == "" || user.Username == "" || user.Email == "" {
		return User{}, errors.New("invalid request, missing required fields")
	}

	return CreateUser(user.Username, user.Password, user.Email)
}

/*
CreateUser creates a new user with the given username, password, and email.
If any of the required fields are missing, it will return an error.
Otherwise, it will return the created user.
*/
func CreateUser(userName, password, email string) (User, error) {
	if userName == "" || password == "" || email == "" {
		return User{}, errors.New("invalid request, missing required fields")
	}

	return User{
		ID:          randStringRunes(10),
		Username:    userName,
		Password:    password,
		Email:       email,
		CreatedDate: time.Now().String(),
		UpdatedDate: time.Now().String(),
	}, nil
}

// randStringRunes generates a random string of length n.
func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
