package main

import (
	"fmt"
	"password-logger/account"
	"password-logger/files"
	"password-logger/output"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	fmt.Println("***Password Logger***")
	vault := account.NewVault(files.NewJsonDb("data.json"))

menuloop:
	for {
		displayMenu()

		choice := menuChoice()

		menuFunc := menu[choice]

		if menuFunc == nil {
			break menuloop
		}

		menuFunc(vault)
	}
}

func displayMenu() {
	fmt.Println("1. Create Account")
	fmt.Println("2. Find Account by URL")
	fmt.Println("3. Find Account by Login")
	fmt.Println("4. Delete Account")
	fmt.Println("5. Exit")
}

func menuChoice() string {
	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	return choice
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scan(&result)
	return result
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Enter login: ")
	password := promptData("Enter password: ")
	url := promptData("Enter URL: ")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Println("could not create an account")
		return
	}

	vault.AddAccount(*myAccount)
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Enter URL: ")

	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})

	if len(accounts) == 0 {
		color.Red("No accounts found")
	}

	for _, acc := range accounts {
		acc.OutputAccountInfo()
	}
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Enter Login: ")

	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})

	if len(accounts) == 0 {
		color.Red("No accounts found")
	}

	for _, acc := range accounts {
		acc.OutputAccountInfo()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Enter URL: ")

	isDeleted := vault.DeleteAccountByUrl(url)

	if isDeleted {
		color.Green("Account deleted")
	} else {
		output.PrintError("Account not found")
	}
}
