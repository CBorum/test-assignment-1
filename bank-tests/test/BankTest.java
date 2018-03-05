import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

class BankTest {

    @Test
    void withdraw() {
        Account a = new Account(123);
        Bank b = new Bank(1000000);
        b.withdraw(a, 123);
        assertEquals(0, a.getBalance());
    }

    @Test
    void deposit() {
        Account a = new Account(123);
        Bank b = new Bank(1000000);
        b.deposit(a, 123);
        assertEquals(246, a.getBalance());
    }

    @Test
    void addAccount() {
        Account a = new Account(123);
        Bank b = new Bank(1000000);
        b.addAccount(a);
        assertEquals(1, b.getAccounts().size());
    }
}