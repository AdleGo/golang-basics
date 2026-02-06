package account

import (
	"encoding/json"
	"password-logger/files"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")

	if err != nil {
		return &Vault{
			Accounts: []Account{},
			UpdateAt: time.Now(),
		}
	}

	var vault Vault

	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red(err.Error())
		return &Vault{
			Accounts: []Account{},
			UpdateAt: time.Now(),
		}
	}

	return &vault
}

func (v *Vault) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, acc := range v.Accounts {
		isMatched := strings.Contains(acc.Url, url)

		if isMatched {
			accounts = append(accounts, acc)
		}
	}

	return accounts
}

func (v *Vault) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false

	for _, acc := range v.Accounts {
		isMatched := strings.Contains(acc.Url, url)

		if isMatched {
			accounts = append(accounts, acc)
			continue
		}
		isDeleted = true
	}

	v.Accounts = accounts
	v.UpdateAt = time.Now()
	data, err := v.ToBytes()

	if err != nil {
		color.Red("could not transform file data.json")
	}
	files.WriteFile(data, "data.json")
	return isDeleted
}

func (v *Vault) AddAccount(acc Account) {
	v.Accounts = append(v.Accounts, acc)
	v.UpdateAt = time.Now()
	data, err := v.ToBytes()

	if err != nil {
		color.Red("could not transform file data.json")
	}
	files.WriteFile(data, "data.json")
}

func (v *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(v)

	if err != nil {
		return nil, err
	}

	return file, nil
}
