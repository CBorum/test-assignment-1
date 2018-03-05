package bank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAccount(t *testing.T) {
	b := &boremiumBank{Balance: 1000000, Accounts: make([]*account, 0)}
	a := &account{0, 1000, 0, "toilet box"}
	b.addAccount(a)

	assert.Equal(t, a, b.Accounts[0])
}

func TestWithdraw(t *testing.T) {
	b := &boremiumBank{Balance: 1000000, Accounts: make([]*account, 0)}
	a := &account{0, 1000, 0, "toilet box"}
	b.addAccount(a)

	balance, err := b.withdraw(a, 100)
	assert.Nil(t, err)
	assert.Equal(t, 900.0, a.Balance)
	assert.Equal(t, 900.0, balance)
}

func TestDeposit(t *testing.T) {
	b := &boremiumBank{Balance: 1000000, Accounts: make([]*account, 0)}
	a := &account{0, 1000, 0, "toilet box"}
	b.addAccount(a)

	balance := b.deposit(a, 100)
	assert.Equal(t, 1100.0, a.Balance)
	assert.Equal(t, 1100.0, balance)
}

func TestMakeLoan(t *testing.T) {
	b := &boremiumBank{Balance: 1000000, Accounts: make([]*account, 0)}
	a := &account{1000, 1000, 0, "toilet box"}
	b.addAccount(a)

	err := b.makeLoan(a, 500)
	assert.Nil(t, err)
	assert.Equal(t, 500.0, a.Loan)
}

func TestPayLoan(t *testing.T) {
	b := &boremiumBank{Balance: 1000000, Accounts: make([]*account, 0)}
	a := &account{1000, 1000, 0, "toilet box"}
	b.addAccount(a)

	err := b.makeLoan(a, 500)
	balance := b.payLoan(a, 500)
	assert.Nil(t, err)
	assert.Equal(t, 0.0, balance)
}

func TestChangeAccountName(t *testing.T) {
	b := &boremiumBank{Balance: 1000000, Accounts: make([]*account, 0)}
	a := &account{1000, 1000, 0, "toilet box"}
	b.addAccount(a)

	b.changeAccountName(a, "abc")
	assert.Equal(t, "abc", a.AccountName)
}

func TestReadAccounts(t *testing.T) {
	b := &boremiumBank{Balance: 1000000, Accounts: make([]*account, 0)}
	b.readAccounts()

}

func TestGetLoaners(t *testing.T) {
	b := &boremiumBank{Balance: 1000000, Accounts: make([]*account, 0)}
	a1 := &account{1000, 1000, 100, "toilet box"}
	a2 := &account{1000, 1000, 102, "toilet box"}
	a3 := &account{1000, 1000, 0, "toilet box"}
	b.addAccount(a1)
	b.addAccount(a2)
	b.addAccount(a3)

	printBank(b)

	accs := b.getLoaners()
	assert.Equal(t, 2, len(accs))
}

func TestGetRichest(t *testing.T) {
	b := &boremiumBank{Balance: 1000000, Accounts: make([]*account, 0)}
	a := &account{1000, 1000, 0, "toilet box"}
	b.addAccount(a)

	b.getRichest()
}

func TestWriteAccountsJSON(t *testing.T) {
	b := &boremiumBank{Balance: 1000000, Accounts: make([]*account, 0)}
	a := &account{1000, 1000, 0, "toilet box"}
	b.addAccount(a)

	err := b.writeAccountsJSON()
	assert.Nil(t, err)
}

func TestRemoveAccount(t *testing.T) {
	b := &boremiumBank{Balance: 1000000, Accounts: make([]*account, 0)}
	a := &account{1000, 1000, 0, "toilet box"}
	b.addAccount(a)

	b.removeAccount(a.AccountName)
	assert.Equal(t, 0, len(b.Accounts))
}

func TestTotalFunds(t *testing.T) {
	b := &boremiumBank{Balance: 1000000, Accounts: make([]*account, 0)}
	a := &account{1000, 1000, 0, "toilet box"}
	b.addAccount(a)
}
