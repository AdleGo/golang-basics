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
	login    string
	password string
	url      string
}

type AccountWithTimestamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (a AccountWithTimestamp) OutputAccountInfo() {
	color.Cyan(a.login)
	color.Cyan(a.password)
	color.Cyan(a.url)
}

func (a *Account) generatePassword(n int) {
	result := make([]rune, n)

	for i := range result {
		result[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	a.password = string(result)
}

func NewAccount(login, password, urlString string) (*AccountWithTimestamp, error) {
	if login == "" {
		return nil, errors.New("Login cannot be empty")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("Invalud URL")
	}

	newAcc := &AccountWithTimestamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			login:    login,
			url:      urlString,
			password: password,
		},
	}

	if password == "" {
		newAcc.Account.generatePassword(12)
	}

	return newAcc, nil
}
