package bank

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type bank interface {
	// adds an account
	addAccount(*account)

	// withdraw amount, returns balance or err
	withdraw(*account, float64) (float64, error)

	// deposit amount, returns balance
	deposit(*account, float64) float64

	// creates a loan for the account, returns error if loan cannot be made
	makeLoan(*account, float64) error

	// pays amount from the account towards the loan, returns the remaining loan amount
	payLoan(*account, float64) float64

	// changes name for account
	changeAccountName(*account, string)

	// reads all accounts from accounts.json
	readAccounts()

	// returns all accounts
	getAccounts() []*account

	// returns accounts that have a loan
	getLoaners() []*account

	// return the account with most money
	getRichest() *account

	// returns all the accounts in json byte slice
	writeAccountsJSON() error

	// removes an account by name
	removeAccount(string)
}

type account struct {
	CreditAllowed float64
	Balance       float64
	Loan          float64
	AccountName   string
}

type boremiumBank struct {
	Balance  float64
	Accounts []*account
}

func (b *boremiumBank) withdraw(a *account, amount float64) (float64, error) {
	if a.Balance-amount < a.CreditAllowed {
		return a.Balance, errors.New("not allowed")
	}
	a.Balance -= amount
	return a.Balance, nil
}

func (b *boremiumBank) addAccount(a *account) {
	b.Accounts = append(b.Accounts, a)
}

func (b *boremiumBank) deposit(a *account, amount float64) float64 {
	a.Balance += amount
	return a.Balance
}

func (b *boremiumBank) makeLoan(a *account, amount float64) error {
	if a.Loan != 0 {
		return errors.New("not allowed")
	}
	a.Loan = amount
	b.Balance -= amount
	return nil
}

func (b *boremiumBank) payLoan(a *account, amount float64) float64 {
	a.Loan -= amount
	b.Balance += amount
	return a.Loan
}

func (b *boremiumBank) changeAccountName(a *account, name string) {
	a.AccountName = name
}

func (b *boremiumBank) readAccounts() {
	file, _ := os.Open("./accounts.json")
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, b.Accounts)
}

func (b *boremiumBank) getLoaners() []*account {
	var loaners []*account
	for _, v := range b.Accounts {
		if v.Loan > 0 {
			loaners = append(loaners, v)
		}
	}
	return loaners
}

func (b *boremiumBank) getRichest() *account {
	var richest *account
	for _, v := range b.Accounts {
		if richest == nil || v.Loan > 0 {
			richest = v
		}
	}
	return richest
}

func (b *boremiumBank) writeAccountsJSON() error {
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

func (b *boremiumBank) removeAccount(name string) {
	for k, v := range b.Accounts {
		if v.AccountName == name {
			b.Accounts = append(b.Accounts[:k], b.Accounts[:k+1]...)
		}
	}
}

func (b *boremiumBank) getAccounts() []*account {
	return b.Accounts
}

func totalFunds(b bank) (sum float64) {
	for _, v := range b.getAccounts() {
		sum += v.Balance
	}
	return
}

func printBank(b bank) {
	for _, v := range b.getAccounts() {
		log.Println(v)
	}
}
