package main

import (
	"fmt"
	"time"
)

/*
===========================================
  Go 100 Challenge
  Problem: #016
  Level: 🟡 Medium
===========================================

Problem: Bank Transaction System

Topic:
- struct
- slice
- functions
- transaction history
- error handling

Industry Use:
Banking Systems / Digital Wallets /
Payment Gateways / Fintech Apps

Rules (English):
- Manage bank account with full transaction history
- Each account has: ID, Owner, Balance
- Each transaction has: Type, Amount, Date, Balance After
- Transaction types: Deposit, Withdraw, Transfer
- User can:
    * Deposit money
    * Withdraw money
    * Transfer to another account
    * View transaction history
    * View balance
    * Exit
- Withdraw more than balance  => "Insufficient balance!"
- Transfer to same account    => "Cannot transfer to same account!"
- Amount <= 0                 => "Amount must be positive!"
- Account not found           => "Account not found!"

Rules (বাংলায়):
- Full transaction history সহ bank account manage করবে
- প্রতিটা account এ থাকবে: ID, Owner, Balance
- প্রতিটা transaction এ থাকবে: Type, Amount, Date, Balance After
- Transaction types: Deposit, Withdraw, Transfer
- User করতে পারবে:
    * টাকা deposit করা
    * টাকা withdraw করা
    * অন্য account এ transfer করা
    * Transaction history দেখা
    * Balance দেখা
    * Exit করা
- Balance এর বেশি withdraw  => "Insufficient balance!"
- Same account এ transfer   => "Cannot transfer to same account!"
- Amount <= 0               => "Amount must be positive!"
- Account না পেলে           => "Account not found!"

Example Run:
  === Bank System ===
  Accounts: ACC001, ACC002

  ACC001 > Choose:
  1. Deposit
  2. Withdraw
  3. Transfer
  4. History
  5. Balance
  6. Exit

  Choose: 1
  Amount: 5000
  ✅ Deposited 5000.00 | Balance: 5000.00

  Choose: 3
  Transfer to (Account ID): ACC002
  Amount: 2000
  ✅ Transferred 2000.00 to ACC002 | Balance: 3000.00

  Choose: 4
  Date         Type       Amount    Balance
  2024-01-15   Deposit    5000.00   5000.00
  2024-01-15   Transfer   2000.00   3000.00
===========================================
*/

func main() {

	var accounts = &[]Account{
		{ID: "111", Owner: "Rakib", Balance: 5000},
		{ID: "222", Owner: "Sumaiya", Balance: 3000},
	}

	MainMenu(accounts)
}

type Account struct {
	ID           string
	Owner        string
	Balance      float64
	Transactions []Transaction
}

type TxnType string

const (
	withdraw TxnType = "Withdraw"
	deposit  TxnType = "Deposit"
	transfer TxnType = "Transfer"
)

type Transaction struct {
	Type         TxnType
	Amount       float64
	Date         time.Time
	BalanceAfter float64
}

func MainMenu(accounts *[]Account) {

	for {
		fmt.Println("┌─────────────────────────────────────┐")
		fmt.Println("│         🏦 BANK SYSTEM              │")
		fmt.Println("│   Available Accounts: ACC001 ACC002 │")
		fmt.Println("└─────────────────────────────────────┘")

		var accountId string
		fmt.Println("Enter Account ID:")
		fmt.Scan(&accountId)

		var foundAccount *Account
		for i := range *accounts {
			if (*accounts)[i].ID == accountId {
				foundAccount = &(*accounts)[i]
				break
			}
		}

		if foundAccount == nil {
			fmt.Println("Account not found!")
			return
		}

		for _, acc := range *accounts {

			if accountId == acc.ID {
				fmt.Println("┌─────────────────────────────────────┐")
				fmt.Printf("│  Welcome, %-26s │\n", acc.Owner)
				fmt.Printf("│  Account: %-26s │\n", acc.ID)
				fmt.Printf("│  Balance: %-26.2f │\n", acc.Balance)
				fmt.Println("┌───────────────────────────────────────────────────────────────┐")
				fmt.Println("│ Type       │ Amount     │ Date       │ Balance After         │")
				fmt.Println("├───────────────────────────────────────────────────────────────┤")

				for _, t := range acc.Transactions {
					fmt.Printf(
						"│ %-10s │ %10.2f │ %-10s │ %18.2f │\n",
						t.Type,
						t.Amount,
						t.Date.Format("2006-01-02"),
						t.BalanceAfter,
					)
				}

				fmt.Println("└───────────────────────────────────────────────────────────────┘")

				fmt.Println("┌──────────────────────────────────────────────┐")
				fmt.Println("│              🏦 BANK DASHBOARD               │")
				fmt.Println("├──────────────────────────────────────────────┤")
				fmt.Println("│                                              │")
				fmt.Println("│  [1] 💰 Deposit                              │")
				fmt.Println("│  [2] 💸 Withdraw                             │")
				fmt.Println("│  [3] 🔄 Transfer                             │")
				fmt.Println("│  [4] 📜 Transaction History                  │")
				fmt.Println("│  [5] 📊 Check Balance                        │")
				fmt.Println("│                                              │")
				fmt.Println("│  [0] 🚪 Exit                                 │")
				fmt.Println("│                                              │")
				fmt.Println("└──────────────────────────────────────────────┘")

				var choose int
				fmt.Print("\n👉 Choose an option: ")

				fmt.Scan(&choose)

				switch choose {
				case 1:
					handleDeposit(accounts, accountId)
				case 2:
					handleWithdraw(accounts, accountId)
				case 0:
					fmt.Println("Good bye")
					return
				}
			}

		}
	}

}

// handle deposit money and history
func handleDeposit(accounts *[]Account, accountId string) {
	fmt.Println("┌─────────────────────────────────────┐")
	fmt.Println("│            💰 DEPOSIT               │")
	fmt.Println("└─────────────────────────────────────┘")

	var depositAmount float64
	fmt.Println("Enter Amount: ")

	fmt.Scan(&depositAmount)

	if _, err := fmt.Scan(&depositAmount); err != nil {
		fmt.Println("Invalid input!")
		return
	}

	if depositAmount <= 0 {
		fmt.Println("Deposit amount must be greater than 0!")
		return
	}

	for i := range *accounts {
		if (*accounts)[i].ID == accountId {
			(*accounts)[i].Balance += depositAmount

			(*accounts)[i].Transactions = append(
				(*accounts)[i].Transactions,
				Transaction{
					Type:         deposit,
					Amount:       depositAmount,
					Date:         time.Now(),
					BalanceAfter: (*accounts)[i].Balance,
				},
			)

			fmt.Printf("✅ Deposited: %.2f\n", depositAmount)
			fmt.Printf("💳 New Balance: %.2f\n", (*accounts)[i].Balance)
			return
		}

	}

	fmt.Println("Account not found!")

}

// handle withdraw money and create history
func handleWithdraw(accounts *[]Account, accountId string) {
	fmt.Println("┌─────────────────────────────────────┐")
	fmt.Println("│           💸 WITHDRAW               │")
	fmt.Println("└─────────────────────────────────────┘")

	var withdrawAmount float64
	fmt.Println("Enter Amount: ")

	_, err := fmt.Scan(&withdrawAmount)
	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	if withdrawAmount <= 0 {
		fmt.Println("Deposit amount must be greater than 0!")
		return
	}

	for i := range *accounts {

		if withdrawAmount > (*accounts)[i].Balance {
			fmt.Println("❌ Insufficient balance!")
			fmt.Printf("💳 Current Balance: %.2f\n", (*accounts)[i].Balance)
			return
		}

		if (*accounts)[i].ID == accountId {
			(*accounts)[i].Balance -= withdrawAmount

			(*accounts)[i].Transactions = append(
				(*accounts)[i].Transactions,
				Transaction{
					Type:         withdraw,
					Amount:       withdrawAmount,
					Date:         time.Now(),
					BalanceAfter: (*accounts)[i].Balance,
				},
			)

			fmt.Printf("✅ Withdraw succesfull: %.2f\n", withdrawAmount)
			fmt.Printf("💳 New Balance: %.2f\n", (*accounts)[i].Balance)

			return
		}

	}

	fmt.Println("Account not found!")

}

func handleTransfer(accounts *[]Account, fromAccountID string) {
	fmt.Println("┌─────────────────────────────────────┐")
	fmt.Println("│           💸 TRANSFER               │")
	fmt.Println("└─────────────────────────────────────┘")

	var toAccountID string
	var amount float64

	fmt.Print("Transfer to Account ID: ")
	fmt.Scan(&toAccountID)

	fmt.Print("Amount: ")
	if _, err := fmt.Scan(&amount); err != nil {
		fmt.Println("❌ Invalid input!")
		return
	}

	if amount <= 0 {
		fmt.Println("❌ Transfer amount must be greater than 0!")
		return
	}

	var fromAcc, toAcc *Account

	// find both accounts
	for i := range *accounts {
		if (*accounts)[i].ID == fromAccountID {
			fromAcc = &(*accounts)[i]
		}
		if (*accounts)[i].ID == toAccountID {
			toAcc = &(*accounts)[i]
		}
	}

	if fromAcc == nil {
		fmt.Println("❌ Sender account not found!")
		return
	}

	if toAcc == nil {
		fmt.Println("❌ Receiver account not found!")
		return
	}

	if fromAcc.Balance < amount {
		fmt.Println("❌ Insufficient balance!")
		fmt.Printf("💳 Current Balance: %.2f\n", fromAcc.Balance)
		return
	}

	// tran logic

	fromAcc.Balance -= amount
	toAcc.Balance += amount

	now := time.Now()

	// sender transaction
	fromAcc.Transactions = append(fromAcc.Transactions, Transaction{
		Type:         transfer,
		Amount:       -amount,
		Date:         now,
		BalanceAfter: fromAcc.Balance,
	})

	// receiver transaction
	toAcc.Transactions = append(toAcc.Transactions, Transaction{
		Type:         deposit,
		Amount:       amount,
		Date:         now,
		BalanceAfter: toAcc.Balance,
	})

	fmt.Printf("✅ Transferred: %.2f → %s (%s)\n", amount, toAcc.ID, toAcc.Owner)
	fmt.Printf("💳 Your New Balance: %.2f\n", fromAcc.Balance)
}

func handleTransactionHistory(accounts *[]Account, accountId string) {
	for _, acc := range *accounts {
		if acc.ID == accountId {

			fmt.Println("┌───────────────────────────────────────────────────────────────┐")
			fmt.Println("│                    📜 TRANSACTION HISTORY                     │")
			fmt.Println("├───────────────────────────────────────────────────────────────┤")
			fmt.Println("│ Type       │ Amount     │ Date       │ Balance After         │")
			fmt.Println("├───────────────────────────────────────────────────────────────┤")

			if len(acc.Transactions) == 0 {
				fmt.Println("│                No transactions found                          │")
				fmt.Println("└───────────────────────────────────────────────────────────────┘")
				return
			}

			for _, t := range acc.Transactions {
				fmt.Printf(
					"│ %-10s │ %10.2f │ %-10s │ %18.2f │\n",
					t.Type,
					t.Amount,
					t.Date.Format("2006-01-02"),
					t.BalanceAfter,
				)
			}

			fmt.Println("└───────────────────────────────────────────────────────────────┘")
			return
		}
	}

	fmt.Println("❌ Account not found!")
}

func handleCheckBalance(accounts *[]Account, accountId string) {
	for _, acc := range *accounts {
		if acc.ID == accountId {

			fmt.Println("┌─────────────────────────────────────┐")
			fmt.Println("│           💳 ACCOUNT INFO           │")
			fmt.Println("├─────────────────────────────────────┤")
			fmt.Printf("│ Owner   : %-24s │\n", acc.Owner)
			fmt.Printf("│ Account : %-24s │\n", acc.ID)
			fmt.Printf("│ Balance : %-24.2f │\n", acc.Balance)
			fmt.Println("└─────────────────────────────────────┘")

			return
		}
	}

	fmt.Println("❌ Account not found!")
}
