package main

/*
===========================================
  Go 100 Challenge
  Problem: #006
  Level: 🟢 Small
===========================================

Problem: Bank Account System

Topic:
- struct
- pointer receivers (important!)
- methods
- error handling

Industry Use:
Banking Systems / Digital Wallets /
Payment Gateways / Fintech Apps

Rules:
- Create a BankAccount struct (Owner, Balance)
- Deposit(amount)  => increases balance
- Withdraw(amount) => decreases balance
- Handle the following errors:
    * Deposit/Withdraw amount <= 0  => "Amount must be positive"
    * Withdraw when balance is low  => "Insufficient balance"
- GetBalance() method => returns current balance
- Print statement after every operation

Example Output:
  Account Owner: Rakib
  ================================
  Deposited:  5000.00  | Balance: 5000.00
  Deposited:  3000.00  | Balance: 8000.00
  Withdrew:   2000.00  | Balance: 6000.00
  ERROR: Insufficient balance
  ERROR: Amount must be positive
  ================================
  Final Balance: 6000.00
===========================================
*/

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {

	if amount <= 0 {
		fmt.Println("Amount must be positive")
		return
	}
	b.Balance += amount // b.Balance = b.Balance + amount
	fmt.Printf("Deposited: %.2f | Balance: %.2f\n", amount, b.Balance)
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Invalid withdraw amount")
		return
	}
	if b.Balance <= 0 || amount > b.Balance {
		fmt.Println("Insufficient balance")
		return
	}
	b.Balance -= amount // b.Balance = b.Balance - amount
	fmt.Printf("Withdrew: %.2f | Balance: %.2f\n", amount, b.Balance)

}

func (b *BankAccount) GetBalance() float64 {
	return b.Balance
}

func main() {
	account := BankAccount{Owner: "Rakib", Balance: 0}

	fmt.Printf("Account Owner: %s\n", account.Owner)
	fmt.Println("================================")

	account.Deposit(5000)
	account.Withdraw(69999)

	fmt.Println("================================")
	fmt.Printf("Final Balance: %.2f\n", account.GetBalance())

}
