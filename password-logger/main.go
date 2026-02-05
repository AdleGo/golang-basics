package main

import (
	"fmt"
	"password-logger/account"
	"password-logger/files"
)

func main() {
	files.ReadFile()
	files.WriteFile()
	login := promptData("Enter login: ")
	password := promptData("Enter password: ")
	url := promptData("Enter URL: ")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		panic("Something went wrong")
	}

	myAccount.OutputAccountInfo()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scan(&result)
	return result
}
