package bank

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAccount(t *testing.T) {
	b := &BoremiumBank{Balance: 1000000, Accounts: make([]*Account, 0)}
	a := &Account{0, 1000, 0, "toilet box"}
	b.AddAccount(a)

	assert.Equal(t, a, b.Accounts[0])
}

func TestWithdraw(t *testing.T) {
	b := &BoremiumBank{Balance: 1000000, Accounts: make([]*Account, 0)}
	a := &Account{0, 1000, 0, "toilet box"}
	b.AddAccount(a)

	balance, err := b.Withdraw(a, 100)
	assert.Nil(t, err)
	assert.Equal(t, 900.0, a.Balance)
	assert.Equal(t, 900.0, balance)
}

func TestDeposit(t *testing.T) {
	b := &BoremiumBank{Balance: 1000000, Accounts: make([]*Account, 0)}
	a := &Account{0, 1000, 0, "toilet box"}
	b.AddAccount(a)

	balance := b.Deposit(a, 100)
	assert.Equal(t, 1100.0, a.Balance)
	assert.Equal(t, 1100.0, balance)
}

func TestMakeLoan(t *testing.T) {
	b := &BoremiumBank{Balance: 1000000, Accounts: make([]*Account, 0)}
	a := &Account{1000, 1000, 0, "toilet box"}
	b.AddAccount(a)

	err := b.MakeLoan(a, 500)
	assert.Nil(t, err)
	assert.Equal(t, 500.0, a.Loan)
}

func TestPayLoan(t *testing.T) {
	b := &BoremiumBank{Balance: 1000000, Accounts: make([]*Account, 0)}
	a := &Account{1000, 1000, 0, "toilet box"}
	b.AddAccount(a)

	err := b.MakeLoan(a, 500)
	balance := b.PayLoan(a, 500)
	assert.Nil(t, err)
	assert.Equal(t, 0.0, balance)
}

func TestChangeAccountName(t *testing.T) {
	b := &BoremiumBank{Balance: 1000000, Accounts: make([]*Account, 0)}
	a := &Account{1000, 1000, 0, "toilet box"}
	b.AddAccount(a)

	b.ChangeAccountName(a, "abc")
	assert.Equal(t, "abc", a.AccountName)
}

func TestReadAccounts(t *testing.T) {
	b := &BoremiumBank{Balance: 1000000, Accounts: make([]*Account, 0)}
	b.ReadAccounts()
}

func TestGetLoaners(t *testing.T) {
	b := &BoremiumBank{Balance: 1000000, Accounts: make([]*Account, 0)}
	a1 := &Account{1000, 1000, 100, "toilet box"}
	a2 := &Account{1000, 1000, 102, "toilet box"}
	a3 := &Account{1000, 1000, 0, "toilet box"}
	b.AddAccount(a1)
	b.AddAccount(a2)
	b.AddAccount(a3)

	PrintBank(b)

	accs := b.GetLoaners()
	assert.Equal(t, 2, len(accs))
}

func TestGetRichest(t *testing.T) {
	b := &BoremiumBank{Balance: 1000000, Accounts: make([]*Account, 0)}
	a := &Account{1000, 1000, 0, "toilet box"}
	b.AddAccount(a)

	b.GetRichest()
}

func TestWriteAccountsJSON(t *testing.T) {
	b := &BoremiumBank{Balance: 1000000, Accounts: make([]*Account, 0)}
	a := &Account{1000, 1000, 0, "toilet box"}
	b.AddAccount(a)

	err := b.WriteAccountsJSON()
	assert.Nil(t, err)
}

func TestRemoveAccount(t *testing.T) {
	b := &BoremiumBank{Balance: 1000000, Accounts: make([]*Account, 0)}
	a := &Account{1000, 1000, 0, "toilet box"}
	b.AddAccount(a)

	b.RemoveAccount(a.AccountName)
	log.Println(b.Accounts)
	assert.Equal(t, 0, len(b.Accounts))
	log.Println(b.Accounts)
}

func TestTotalFunds(t *testing.T) {
	b := &BoremiumBank{Balance: 1000000, Accounts: make([]*Account, 0)}
	a := &Account{1000, 1000, 0, "toilet box"}
	b.AddAccount(a)
}
