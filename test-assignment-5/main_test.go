package banktests

import (
	"testing"

	bank "github.com/cborum/test-assignments/test-assignment-4"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestAddAccount(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	a := &bank.Account{0, 1000, 0, "Harry Turnado"}
	b.AddAccount(a)

	then.AssertThat(t, len(b.Accounts), is.EqualTo(1))
}

func TestWithdraw(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	a := &bank.Account{0, 1000, 0, "Harry Turnado"}
	b.AddAccount(a)

	balance, err := b.Withdraw(a, 100)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, a.Balance, is.EqualTo(900.0))
	then.AssertThat(t, balance, is.EqualTo(900.0))
}

func TestDeposit(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	a := &bank.Account{0, 1000, 0, "Harry Turnado"}
	b.AddAccount(a)

	balance := b.Deposit(a, 100)
	then.AssertThat(t, a.Balance, is.EqualTo(1100.0))
	then.AssertThat(t, balance, is.EqualTo(1100.0))
}

func TestMakeLoan(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	a := &bank.Account{1000, 1000, 0, "Harry Turnado"}
	b.AddAccount(a)

	err := b.MakeLoan(a, 500)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, a.Loan, is.EqualTo(500.0))
}

func TestPayLoan(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	a := &bank.Account{1000, 1000, 0, "Harry Turnado"}
	b.AddAccount(a)

	err := b.MakeLoan(a, 500)
	balance := b.PayLoan(a, 500)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, balance, is.EqualTo(0.0))
}

func TestChangeAccountName(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	a := &bank.Account{1000, 1000, 0, "Harry Turnado"}
	b.AddAccount(a)

	b.ChangeAccountName(a, "abc")
	then.AssertThat(t, a.AccountName, is.EqualTo("abc"))
}

func TestReadAccounts(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	b.ReadAccounts()
}

func TestGetLoaners(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	a1 := &bank.Account{1000, 1000, 100, "Harry Turnado"}
	a2 := &bank.Account{1000, 1000, 102, "Harry Turnado"}
	a3 := &bank.Account{1000, 1000, 0, "Harry Turnado"}
	b.AddAccount(a1)
	b.AddAccount(a2)
	b.AddAccount(a3)

	bank.PrintBank(b)

	accs := b.GetLoaners()
	then.AssertThat(t, len(accs), is.EqualTo(2))
}

func TestGetRichest(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	a := &bank.Account{1000, 1000, 0, "Harry Turnado"}
	b.AddAccount(a)

	richest := b.GetRichest()
	then.AssertThat(t, richest, is.EqualTo(a))
}

func TestWriteAccountsJSON(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	a := &bank.Account{1000, 1000, 0, "Harry Turnado"}
	b.AddAccount(a)

	err := b.WriteAccountsJSON()
	then.AssertThat(t, err, is.Nil())
}

func TestRemoveAccount(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	a := &bank.Account{1000, 1000, 0, "Harry Turnado"}
	b.AddAccount(a)

	b.RemoveAccount(a.AccountName)
	then.AssertThat(t, len(b.Accounts), is.EqualTo(0))
}

func TestTotalFunds(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	a := &bank.Account{1000, 1000, 0, "Harry Turnado"}
	b.AddAccount(a)
}

var withdrawtests = []struct {
	balance  float64
	withdraw float64
	result   float64
}{
	{110, 10, 100},
	{220, 10, 210},
	{330, 10, 320},
	{440, 10, 430},
	{550, 10, 540},
}

func TestFlagParser(t *testing.T) {
	b := &bank.BoremiumBank{Balance: 1000000, Accounts: make([]*bank.Account, 0)}
	for _, tt := range withdrawtests {
		account := &bank.Account{0, tt.balance, 0, "Harry Turnado"}
		b.Withdraw(account, tt.withdraw)
		if account.Balance != tt.result {
			t.Errorf("Error: account balance should be %.2f, but was %.2f", account.Balance, tt.result)
		}
	}
}
