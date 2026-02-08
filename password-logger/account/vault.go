package account

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func NewVault(db Db) *VaultWithDb {
	file, err := db.Read()

	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db: db,
		}
	}

	var vault Vault

	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red(err.Error())

		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db: db,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}

func (v *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, acc := range v.Accounts {
		isMatched := checker(acc, str)

		if isMatched {
			accounts = append(accounts, acc)
		}
	}

	return accounts
}

func (v *VaultWithDb) DeleteAccountByUrl(url string) bool {
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
	data, err := v.Vault.ToBytes()

	if err != nil {
		color.Red("could not transform file data.json")
	}

	v.db.Write(data)

	return isDeleted
}

func (v *VaultWithDb) AddAccount(acc Account) {
	v.Accounts = append(v.Accounts, acc)
	v.UpdateAt = time.Now()
	data, err := v.Vault.ToBytes()

	if err != nil {
		color.Red("could not transform file data.json")
	}

	v.db.Write(data)
}

func (v *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(v)

	if err != nil {
		return nil, err
	}

	return file, nil
}
