import java.util.ArrayList;
import java.util.List;

public class Bank implements iBank {
    ArrayList<Account> accounts = new ArrayList<>();
    private double balance;

    public Bank(double balance) {
        this.balance = balance;
    }

    @Override
    public double withdraw(Account a, double value) {
        a.setBalance(a.getBalance() - value);
        return a.getBalance();
    }

    @Override
    public double deposit(Account a, double value) {
        a.setBalance(a.getBalance() + value);
        return a.getBalance();
    }

    @Override
    public void addAccount(Account a) {
        accounts.add(a);
    }

    @Override
    public ArrayList<Account> getAccounts() {
        return accounts;
    }
}
