package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimestamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func (a accountWithTimestamp) outputPassword() {
	fmt.Println(a.login, a.password, a.url)
}

func (a *account) generatePassword(n int) {
	result := make([]rune, n)

	for i := range result {
		result[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	a.password = string(result)
}

func newAccount(login, password, urlString string) (*accountWithTimestamp, error) {
	if login == "" {
		return nil, errors.New("Login cannot be empty")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("Invalud URL")
	}

	newAcc := &accountWithTimestamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			login:    login,
			url:      urlString,
			password: password,
		},
	}

	if password == "" {
		newAcc.account.generatePassword(12)
	}

	return newAcc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQESTUVWXYZ1234567890-*!")

func main() {
	login := promptData("Enter login: ")
	password := promptData("Enter password: ")
	url := promptData("Enter URL: ")

	myAccount, err := newAccount(login, password, url)

	if err != nil {
		panic("Something went wrong")
	}

	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scan(&result)
	return result
}
