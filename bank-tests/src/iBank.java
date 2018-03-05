import java.util.ArrayList;

public interface iBank {
    double withdraw(Account a, double value);
    double deposit(Account a, double value);
    void addAccount(Account a);
    ArrayList<Account> getAccounts();
}
