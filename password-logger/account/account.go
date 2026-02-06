package account

import (
	"errors"
	"math/rand"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQESTUVWXYZ1234567890-*!")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (a Account) OutputAccountInfo() {
	color.Cyan(a.Login)
	color.Cyan(a.Password)
	color.Cyan(a.Url)
}

func (a *Account) generatePassword(n int) {
	result := make([]rune, n)

	for i := range result {
		result[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	a.Password = string(result)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("Login cannot be empty")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("Invalud URL")
	}

	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Url:       urlString,
		Password:  password,
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}
