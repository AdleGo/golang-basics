package main

import (
	"fmt"
	"password-logger/account"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("***Password Logger***")
	vault := account.NewVault()

menuloop:
	for {
		displayMenu()
		choice := menuChoice()

		switch choice {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		case "4":
			break menuloop
		}
	}
}

func displayMenu() {
	fmt.Println("1. Create Account")
	fmt.Println("2. Find Account")
	fmt.Println("3. Delete Account")
	fmt.Println("4. Exit")
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

func createAccount(vault *account.Vault) {
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

func findAccount(vault *account.Vault) {
	url := promptData("Enter URL: ")

	accounts := vault.FindAccountsByUrl(url)

	if len(accounts) == 0 {
		color.Red("No accounts found")
	}

	for _, acc := range accounts {
		acc.OutputAccountInfo()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Enter URL: ")

	isDeleted := vault.DeleteAccountByUrl(url)

	if isDeleted {
		color.Green("Account deleted")
	} else {
		color.Red("Account not found")
	}
}
