package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
===========================================
  Go 100 Challenge
  Problem: #013
  Level: 🟡 Medium
===========================================

Problem: Contact Book CLI

Topic:
- struct
- map
- functions
- string search
- CRUD operations

Industry Use:
CRM Systems / Phone Book Apps /
Contact Management / Address Book

Rules (English):
- Manage contacts with full CRUD
- Each contact has: Name, Phone, Email
- User can:
    * Add contact
    * Search contact by name
    * Update contact (phone or email)
    * Delete contact by name
    * View all contacts
    * Exit
- Duplicate name    => "Contact already exists!"
- Name not found    => "Contact not found!"
- Invalid phone     => must be 11 digits (numbers only)
- Invalid email     => must contain @ and .

Rules (বাংলা):
- Full CRUD দিয়ে contacts manage করবে
- প্রতিটা contact এ থাকবে: Name, Phone, Email
- User করতে পারবে:
    * Contact add করা
    * নাম দিয়ে contact search করা
    * Contact update করা (phone বা email)
    * নাম দিয়ে contact delete করা
    * সব contacts দেখা
    * Exit করা
- Duplicate নাম      => "Contact already exists!"
- নাম না পেলে       => "Contact not found!"
- Invalid phone     => ১১ digits হতে হবে (শুধু numbers)
- Invalid email     => @ এবং . থাকতে হবে

Example Run:
  === Contact Book ===
  1. Add Contact
  2. Search Contact
  3. Update Contact
  4. Delete Contact
  5. View All Contacts
  6. Exit

  Choose: 1
  Name: Rakib
  Phone: 01712345678
  Email: rakib@gmail.com
  Contact added!

  Choose: 2
  Search by name: Rakib
  Name : Rakib
  Phone: 01712345678
  Email: rakib@gmail.com

  Choose: 3
  Name to update: Rakib
  1. Update Phone
  2. Update Email
  Choose: 1
  New Phone: 01987654321
  Updated successfully!

  Choose: 4
  Name to delete: Rakib
  Contact deleted!

  Choose: 5
  No contacts found!
===========================================
*/

type Contact struct {
	name  string
	phone int
	email string
}

var reader = bufio.NewReader(os.Stdin)

func readInput() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	MainMenu()
}

func MainMenu() {
	contact := &[]Contact{}
	var choose int
	for {
		fmt.Println("\n┌──────────────────────────────────────┐")
		fmt.Println("│        📱 CONTACT BOOK SYSTEM        │")
		fmt.Println("├──────────────────────────────────────┤")
		fmt.Println("│  1. Add New Contact                  │")
		fmt.Println("│  2. Search Contact                   │")
		fmt.Println("│  3. Update Contact                   │")
		fmt.Println("│  4. Delete Contact                   │")
		fmt.Println("│  5. View All Contacts                │")
		fmt.Println("│  6. Exit System                      │")
		fmt.Println("└──────────────────────────────────────┘")
		fmt.Print("👉 Choose an option (1-6): ")

		fmt.Scanln(&choose)

		switch choose {
		case 1:
			createContact(contact)
		case 2:
			searchContact(contact)
		case 3:
			updateContactInfo(contact)
		case 4:
			deleteContact(contact)
		case 5:
			viewAllContacts(contact)
		case 6:
			return

		}
	}
}

//   Choose: 1
//   Name: Rakib
//   Phone: 01712345678
//   Email: rakib@gmail.com
//   Contact added!

func createContact(contact *[]Contact) {

	fmt.Println("Enter Name: ")
	inputName := readInput()

	if len(inputName) == 0 {
		fmt.Println("Invalid input!")
	}

	fmt.Println("Enter Phone: ")
	inputPhone := readInput()

	if len(inputPhone) == 0 {
		fmt.Println("Invalid input!")
	}

	phoneNum, error := strconv.Atoi(inputPhone)

	if error != nil {
		fmt.Println("error", error.Error())
	}

	// if phoneNum != 11 {
	// 	fmt.Println("⚠️ must be 11 digit")
	// 	return
	// }

	fmt.Print("Enter Email: ")
	inputEmail := readInput()

	inputEmail = strings.TrimSpace(inputEmail)

	if len(inputEmail) == 0 {
		fmt.Println("Email is required")
		return
	}

	email := strings.ToLower(inputEmail)

	if !strings.Contains(email, "@") ||
		!strings.Contains(email, ".") {

		fmt.Println("Invalid email format")
		return
	}

	newContact := Contact{
		name:  inputName,
		phone: phoneNum,
		email: email,
	}

	*contact = append(*contact, newContact)

	fmt.Println("ccont", *contact)

	fmt.Println(" Contact added successfully!")

}

//   Choose: 2
//   Search by name: Rakib
//   Name : Rakib
//   Phone: 01712345678
//   Email: rakib@gmail.com

func searchContact(contacts *[]Contact) {
	fmt.Println("Search by name: ")
	input := readInput()

	for index, contact := range *contacts {

		if strings.Contains(
			strings.ToLower(contact.name),
			strings.ToLower(input)) {
			fmt.Println("index:", index+1)
			fmt.Printf("│ Name   : %-34s │\n", contact.name)
			fmt.Printf("│ Phone  : %-34d │\n", contact.phone)
			fmt.Printf("│ Email  : %-34s │\n", contact.email)

		}
	}
}

//   Choose: 3
//   Name to update: Rakib
//   1. Update Phone
//   2. Update Email
//   Choose: 1
//   New Phone: 01987654321
//   Updated successfully!

func updateContactInfo(contacts *[]Contact) {
	if len(*contacts) == 0 {
		fmt.Println("📭 Your contact book is empty!")
		return
	}

	fmt.Println("Enter name: ")
	nameInput := readInput()

	isFound := false
	for index := range *contacts {
		if strings.Contains(
			strings.ToLower((*contacts)[index].name),
			strings.ToLower(nameInput),
		) {

			isFound = true
			var choice int

			fmt.Println("1. Update Phone")
			fmt.Println("2. Update Email")
			fmt.Println("3. Home")
			fmt.Println("Choose: ")

			_, err := fmt.Scanln(&choice)

			if err != nil {
				fmt.Println("❌ Invalid menu choice!")
				return
			}
			switch choice {
			case 1:
				updatePhoneNum(&(*contacts)[index].phone)
			case 2:
				updateEmail(&(*contacts)[index].email)
			case 3:
				MainMenu()
			}

		}
	}

	if !isFound {
		fmt.Printf("❌ No contact found matching the name: '%s'\n", nameInput)
	}
}

// update phone number logic
func updatePhoneNum(phone *int) (int, error) {
	fmt.Println("Enter new phone number: ")
	input := readInput()

	if len(input) == 0 {
		fmt.Println("Invalid input")
	}

	convertedInput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number format!")
		return 0, err
	}

	*phone = convertedInput

	fmt.Println("Phone updated successfully!")

	return *phone, nil
}

// update email logic
func updateEmail(email *string) string {
	fmt.Println("Enter new email: ")
	input := readInput()

	if len(input) == 0 {
		fmt.Println("Invalid input")
	}

	*email = input

	fmt.Println("Email updated successfully!")

	return *email
}

//   Choose: 4
//   Name to delete: Rakib
//   Contact deleted!

func deleteContact(contacts *[]Contact) {
	fmt.Println("Enter Name to delete: ")
	input := readInput()

	if len(input) == 0 {
		fmt.Println("Invalid Input")
	}

	isFound := false

	for i := range *contacts {
		if strings.Contains(
			strings.ToLower((*contacts)[i].name),
			strings.ToLower(input),
		) {
			isFound = true
			fmt.Printf("🗑️ Deleting contact: %s...\n", (*contacts)[i].name)

			*contacts = append((*contacts)[:i], (*contacts)[i+1:]...)

			fmt.Println("✅ Contact deleted successfully!")
		}
	}

	if !isFound {
		fmt.Printf("❌ No contact found with the name: '%s'\n", input)
	}
}

func viewAllContacts(contacts *[]Contact) {

	if len(*contacts) == 0 {
		fmt.Println("📭 Your contact book is empty! No contacts to display.")
		return
	}

	fmt.Println("\n=================== 📋 ALL CONTACTS ===================")

	for index, contact := range *contacts {
		fmt.Println("index:", index+1)
		fmt.Println("┌──────────────────────────────────────────────────┐")
		fmt.Printf("│ Name   : %-34s │\n", contact.name)
		fmt.Printf("│ Phone  : %-34d │\n", contact.phone)
		fmt.Printf("│ Email  : %-34s │\n", contact.email)
		fmt.Println("└──────────────────────────────────────────────────┘")
	}

	fmt.Printf("================ Total Contacts: %d ================\n", len(*contacts))
}
