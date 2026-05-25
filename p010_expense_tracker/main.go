package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
===========================================
  Go 100 Challenge
  Problem: #010
  Level: 🟡 Medium
===========================================

Problem: Expense Tracker CLI

Topic:
- struct
- slice
- map
- functions
- date/time (time package)

Industry Use:
Personal Finance Apps / Accounting Software /
Budget Management / Fintech

Rules (English):
- User can add an expense with: Title, Amount, Category
- Categories: Food, Transport, Shopping, Bills, Other
- User can:
    * Add expense
    * View all expenses
    * View total by category
    * View grand total
    * Exit
- Invalid category => "Invalid category!"
- Empty title => "Title cannot be empty!"
- Amount <= 0  => "Amount must be positive!"

Rules (বাংলা):
- User expense add করবে: Title, Amount, Category দিয়ে
- Categories: Food, Transport, Shopping, Bills, Other
- User করতে পারবে:
    * Expense add করা
    * সব expense দেখা
    * Category অনুযায়ী total দেখা
    * Grand total দেখা
    * Exit করা
- Invalid category  => "Invalid category!"
- খালি title       => "Title cannot be empty!"
- Amount <= 0      => "Amount must be positive!"

Example Run:
  === Expense Tracker ===
  1. Add Expense
  2. View All Expenses
  3. View By Category
  4. Grand Total
  5. Exit

  Choose: 1
  Title: Lunch
  Amount: 150
  Category: Food
  Expense added!

  Choose: 3
  Food        : 150.00
  Transport   : 0.00
  Shopping    : 0.00
  Bills       : 0.00
  Other       : 0.00

  Choose: 4
  Grand Total: 150.00
===========================================
*/

type Expense struct {
	Title    string
	Amount   float64
	Category string
}

var reader = bufio.NewReader(os.Stdin)

func readInput() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {

	expense := &[]Expense{}

	var choice int

	for {
		fmt.Println("=== Expense Tracker ===")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View All Expenses")
		fmt.Println("3. View By Category")
		fmt.Println("4. Grand Total")
		fmt.Println("5. Exit")
		fmt.Print("Choose: ")

		// input
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addExpense(expense)
		case 2:
			viewAllExpense(expense)
		// case 3:
		// 	calcCatogry(expense)
		case 3:
			viewMyCategories(expense)
		case 4:
			grandTotal(expense)
		case 5:
			return
		}

	}

}

func addExpense(expense *[]Expense) {

	fmt.Println("Enter your Title:")
	titleInput := readInput()
	if len(titleInput) == 0 {
		fmt.Println("Title cannot be empty!")
		return
	}

	fmt.Println("Enter you amount:")
	amountInput := readInput()

	if len(amountInput) == 0 {
		fmt.Println("Amount cannot be empty!")
		return
	}

	var convertedNum float64
	fmt.Sscan(amountInput, &convertedNum)

	var categoryInput string

	fmt.Println("=== categories ===")
	fmt.Println("1. Food")
	fmt.Println("2.	Transport")
	fmt.Println("3. Shopping")
	fmt.Println("4. Bills")
	fmt.Println("5. Other")
	fmt.Print("Choose: ")

	chooseCategory := readInput()

	switch chooseCategory {
	case "1":
		categoryInput = "Food"
	case "2":
		categoryInput = "Transport"
	case "3":
		categoryInput = "Shopping"
	case "4":
		categoryInput = "Bills"
	case "5":
		categoryInput = "Other"
	default:
		fmt.Println("Invalid category!")
		return
	}

	newExpense := Expense{
		Title:    titleInput,
		Amount:   convertedNum,
		Category: categoryInput,
	}

	*expense = append(*expense, newExpense)

	fmt.Println("Expense added!")

	fmt.Println("appended:", *expense)

}

func viewAllExpense(expenses *[]Expense) {

	if len(*expenses) == 0 {
		fmt.Println("No expenses found!")
		return
	}

	for i, exp := range *expenses {
		fmt.Println("exp:", i+1, exp)
	}
}

func viewMyCategories(expenses *[]Expense) {
	total := make(map[string]float64)

	for _, exp := range *expenses {
		total[exp.Category] += exp.Amount
	}

	fmt.Println("Expense Totals by Category:")
	for category, total := range total {
		fmt.Printf("- %s: %.2f\n", category, total)
	}
}

// func calcCatogry(expenses *[]Expense) {
// 	total := make(map[string]float64)

// 	for _, exp := range *expenses {
// 		total[exp.Category] += exp.Amount
// 	}

// 	fmt.Println("Expense Totals by Category:")
// 	for category, total := range total {
// 		fmt.Printf("- %s: %.2f\n", category, total)
// 	}
// }

// grandTotal calculates the sum of all expenses
func grandTotal(expenses *[]Expense) float64 {
	total := 0.0
	for _, exp := range *expenses {
		total += exp.Amount
	}
	return total
}
