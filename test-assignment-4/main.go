package bank

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Bank interface {
	// adds an account
	AddAccount(*Account)

	// withdraw amount, returns balance or err
	Withdraw(*Account, float64) (float64, error)

	// deposit amount, returns balance
	Deposit(*Account, float64) float64

	// creates a loan for the account, returns error if loan cannot be made
	MakeLoan(*Account, float64) error

	// pays amount from the account towards the loan, returns the remaining loan amount
	PayLoan(*Account, float64) float64

	// changes name for account
	ChangeAccountName(*Account, string)

	// reads all accounts from accounts.json
	ReadAccounts()

	// returns all accounts
	GetAccounts() []*Account

	// returns accounts that have a loan
	GetLoaners() []*Account

	// return the account with most money
	GetRichest() *Account

	// returns all the accounts in json byte slice
	WriteAccountsJSON() error

	// removes an account by name
	RemoveAccount(string)
}

type Account struct {
	CreditAllowed float64
	Balance       float64
	Loan          float64
	AccountName   string
}

type BoremiumBank struct {
	Balance  float64
	Accounts []*Account
}

func (b *BoremiumBank) Withdraw(a *Account, amount float64) (float64, error) {
	if a.Balance-amount < a.CreditAllowed {
		return a.Balance, errors.New("not allowed")
	}
	a.Balance -= amount
	return a.Balance, nil
}

func (b *BoremiumBank) AddAccount(a *Account) {
	b.Accounts = append(b.Accounts, a)
}

func (b *BoremiumBank) Deposit(a *Account, amount float64) float64 {
	a.Balance += amount
	return a.Balance
}

func (b *BoremiumBank) MakeLoan(a *Account, amount float64) error {
	if a.Loan != 0 {
		return errors.New("not allowed")
	}
	a.Loan = amount
	b.Balance -= amount
	return nil
}

func (b *BoremiumBank) PayLoan(a *Account, amount float64) float64 {
	a.Loan -= amount
	b.Balance += amount
	return a.Loan
}

func (b *BoremiumBank) ChangeAccountName(a *Account, name string) {
	a.AccountName = name
}

func (b *BoremiumBank) ReadAccounts() {
	file, _ := os.Open("./accounts.json")
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, b.Accounts)
}

func (b *BoremiumBank) GetLoaners() []*Account {
	var loaners []*Account
	for _, v := range b.Accounts {
		if v.Loan > 0 {
			loaners = append(loaners, v)
		}
	}
	return loaners
}

func (b *BoremiumBank) GetRichest() *Account {
	var richest *Account
	for _, v := range b.Accounts {
		if richest == nil || v.Loan > 0 {
			richest = v
		}
	}
	return richest
}

func (b *BoremiumBank) WriteAccountsJSON() error {
	bytes, err := json.Marshal(b.Accounts)
	if err != nil {
		return err
	}
	file, err := os.OpenFile("./accounts.json", os.O_WRONLY, 0666)
	if err != nil {
		return nil
	}
	_, err = file.Write(bytes)
	return err
}

func (b *BoremiumBank) RemoveAccount(name string) {
	for k, v := range b.Accounts {
		if v.AccountName == name {
			b.Accounts = append(b.Accounts[:k], b.Accounts[:k+1]...)
		}
	}
}

func (b *BoremiumBank) GetAccounts() []*Account {
	return b.Accounts
}

func TotalFunds(b Bank) (sum float64) {
	for _, v := range b.GetAccounts() {
		sum += v.Balance
	}
	return
}

func PrintBank(b Bank) {
	for _, v := range b.GetAccounts() {
		log.Println(v)
	}
}
