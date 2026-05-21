package main

import "fmt"

/*
===========================================
  Go 100 Challenge
  Problem: #002
  Level: 🟢 Small
===========================================

Problem: Failed Login Attempt Tracker

Topic:
- for loop
- break
- if/else
- authentication logic

Industry Use:
Security System / ATM Authentication

Rules:
- Correct PIN = 1234
- User gets maximum 3 attempts
- Correct PIN    => "Access Granted!"
- Wrong PIN      => "Wrong PIN! X attempts remaining"
- After 3 fails  => "Card Blocked!"

Example Run 1 (correct on 2nd try):
  Enter PIN: 0000
  Wrong PIN! 2 attempts remaining
  Enter PIN: 1234
  Access Granted!

Example Run 2 (all wrong):
  Enter PIN: 0000
  Wrong PIN! 2 attempts remaining
  Enter PIN: 1111
  Wrong PIN! 1 attempts remaining
  Enter PIN: 2222
  Card Blocked!
===========================================
*/

func atmSecurity(PIN int, maxAttempts int) {

	var userInput int

	for i := 1; i <= maxAttempts; i++ {

		fmt.Println("Enter PIN:")
		fmt.Scan(&userInput)

		if userInput == PIN {
			fmt.Println("Access Granted!")
			return
		} else {
			remaining := maxAttempts - i

			if remaining > 0 {
				fmt.Printf("Wrong PIN! %d attempts remaining\n", remaining)
			}
		}
	}

	fmt.Println("Card Blocked!")

}

func main() {

	correctPIN := 1234
	maxAttempts := 3
	atmSecurity(correctPIN, maxAttempts)
}
