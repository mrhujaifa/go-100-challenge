package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
===========================================
  Go 100 Challenge
  Problem: #009
  Level: 🟢 Small
===========================================

Problem: Password Validator

Topic:
- strings package
- for range on string
- unicode package
- multiple conditions
- functions

Industry Use:
Auth Systems / User Registration /
Security Tools / API Authentication

Rules (English):
- Validate password with these rules:
    * Minimum 8 characters
    * At least 1 uppercase letter (A-Z)
    * At least 1 lowercase letter (a-z)
    * At least 1 number (0-9)
    * At least 1 special character (!@#$%^&*)
- All rules pass => "Password is strong!"
- Any rule fail  => show which rules failed

Rules (বাংলা):
- Password এই rules দিয়ে validate করবে:
    * সর্বনিম্ন ৮টি character থাকতে হবে
    * কমপক্ষে ১টি uppercase letter (A-Z)
    * কমপক্ষে ১টি lowercase letter (a-z)
    * কমপক্ষে ১টি number (0-9)
    * কমপক্ষে ১টি special character (!@#$%^&*)
- সব rule pass হলে => "Password is strong!"
- Fail হলে => কোন rule fail হয়েছে সেটা দেখাবে

Example Run 1:
  Enter password: Hello1!
  Weak password! Reasons:
  - Minimum 8 characters required

Example Run 2:
  Enter password: hello123!
  Weak password! Reasons:
  - At least 1 uppercase letter required

Example Run 3:
  Enter password: Hello123!
  Password is strong!
===========================================
*/

var reader = bufio.NewReader(os.Stdin)

func readInput() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {

	fmt.Print("Enter password: ")

	pass := readInput()

	isValid, message := passwordValidation(pass)

	fmt.Println("status:", isValid)
	fmt.Println("Error:", message)
}

func passwordValidation(password string) (bool, string) {

	if len(password) < 8 {
		return false, "password must be at least 8 characters long"
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, runeCodePass := range password {

		switch {

		case unicode.IsUpper(runeCodePass):
			hasUpper = true

		case unicode.IsLower(runeCodePass):
			hasLower = true

		case unicode.IsDigit(runeCodePass):
			hasDigit = true

		case unicode.IsPunct(runeCodePass) || unicode.IsSymbol(runeCodePass):
			hasSpecial = true
		}
	}
	errors := []string{}

	if len(password) < 8 {
		errors = append(errors, "- Minimum 8 characters required")
	}
	if !hasUpper {
		errors = append(errors, "- At least 1 uppercase letter required")
	}
	if !hasLower {
		errors = append(errors, "- At least 1 lowercase letter required")
	}
	if !hasDigit {
		errors = append(errors, "- At least 1 number required")
	}
	if !hasSpecial {
		errors = append(errors, "- At least 1 special character required")
	}

	if len(errors) > 0 {
		return false, "Weak password! Reasons:\n" + strings.Join(errors, "\n")
	}

	return true, "Password is strong!"
}
