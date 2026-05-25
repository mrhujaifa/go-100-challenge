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
  Problem: #011
  Level: 🟡 Medium
===========================================

Problem: Library Book Manager

Topic:
- struct
- slice
- functions
- string search
- for loop

Industry Use:
Library Management Systems / Book Store Apps /
Digital Catalogues / Inventory Systems

Rules (English):
- Manage a library with books
- Each book has: Title, Author, Year, IsAvailable
- User can:
    * Add a book
    * Search book by title
    * Borrow a book (IsAvailable = false)
    * Return a book (IsAvailable = true)
    * View all books with availability status
    * Exit
- Borrow unavailable book => "Book is already borrowed!"
- Return available book   => "Book is not borrowed!"
- Search no result        => "No book found!"

Rules (বাংলা):
- একটা library manage করবে books দিয়ে
- প্রতিটা book এ থাকবে: Title, Author, Year, IsAvailable
- User করতে পারবে:
    * Book add করা
    * Title দিয়ে book search করা
    * Book borrow করা (IsAvailable = false)
    * Book return করা (IsAvailable = true)
    * সব books দেখা availability সহ
    * Exit করা
- Already borrowed book borrow করলে => "Book is already borrowed!"
- Available book return করলে        => "Book is not borrowed!"
- Search এ কিছু না পেলে             => "No book found!"

Example Run:
  === Library Manager ===
  1. Add Book
  2. Search Book
  3. Borrow Book
  4. Return Book
  5. View All Books
  6. Exit

  Choose: 1
  Title: The Go Programming Language
  Author: Alan Donovan
  Year: 2015
  Book added!

  Choose: 1
  Title: Clean Code
  Author: Robert Martin
  Year: 2008
  Book added!

  Choose: 5
  [AVAILABLE] The Go Programming Language - Alan Donovan (2015)
  [AVAILABLE] Clean Code - Robert Martin (2008)

  Choose: 2
  Enter title to search: Clean Code
  [AVAILABLE] Clean Code - Robert Martin (2008)

  Choose: 2
  Enter title to search: Harry Potter
  No book found!

  Choose: 3
  Enter book title to borrow: The Go Programming Language
  Book borrowed successfully!

  Choose: 3
  Enter book title to borrow: The Go Programming Language
  Book is already borrowed!

  Choose: 5
  [BORROWED]   The Go Programming Language - Alan Donovan (2015)
  [AVAILABLE]  Clean Code - Robert Martin (2008)

  Choose: 4
  Enter book title to return: The Go Programming Language
  Book returned successfully!

  Choose: 4
  Enter book title to return: The Go Programming Language
  Book is not borrowed!

  Choose: 6
  Goodbye!
===========================================
*/

type Books struct {
	title       string
	author      string
	year        string
	IsAvailable bool
}

var reader = bufio.NewReader(os.Stdin)

func readInput() string {
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func main() {
	books := &[]Books{}

	var choice int
	for {
		fmt.Println("========================================")
		fmt.Println("    ____  ____   ____  _  __")
		fmt.Println("   / __ )/ __ \\ / __ \\| |/ /")
		fmt.Println("  / __  / / / // / / /|   / ")
		fmt.Println(" / /_/ / /_/ // /_/ //   |  ")
		fmt.Println("/_____/\\____/ \\____//_/|_|  ")
		fmt.Println("       LIBRARY MANAGER       ")
		fmt.Println("========================================")

		fmt.Println(" 1. Add Book")
		fmt.Println(" 2. Search Book")
		fmt.Println(" 3. Borrow Book")
		fmt.Println(" 4. Return Book")
		fmt.Println(" 5. View All Books")
		fmt.Println(" 6. Exit")
		fmt.Println("========================================")
		fmt.Println("Enter your choice (1-6):")

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addBook(books)
		case 2:
			searchBooks(books)
		case 3:
			bookTitleToBorrow(books)
		case 4:
			bookTitleToReturn(books)
		case 5:
			viewAllBooks(books)
		case 6:
			return
		}

	}

}

// input validation check logic
func getValidInput(label string, maxLength int, requiredLength int) string {

	for {
		fmt.Print(label)
		input := readInput()

		// empty check
		if len(input) == 0 {
			fmt.Printf("\033[31m✕ Error: Input cannot be empty! Try again.\033[0m\n")
			continue
		}

		// max length check (optional)
		if maxLength > 0 && len(input) > maxLength {
			fmt.Printf(
				"\033[31m✕ Error: Input cannot be more than %d characters! Try again.\033[0m\n",
				maxLength,
			)
			continue
		}

		// required exact length check (optional)
		if requiredLength > 0 && len(input) != requiredLength {
			fmt.Printf(
				"\033[31m✕ Error: Input must be exactly %d characters! Try again.\033[0m\n",
				requiredLength,
			)
			continue
		}

		return input
	}
}

// add book logic
func addBook(books *[]Books) {

	titleInput := getValidInput("Title: ", 0, 0)

	authorInput := getValidInput("Author: ", 4, 0)
	yearInput := getValidInput("Year:", 4, 4)

	newBooks := Books{
		title:       titleInput,
		author:      authorInput,
		year:        yearInput,
		IsAvailable: true,
	}

	*books = append(*books, newBooks)
	fmt.Println("\n┌──────────────────────────────────────┐")
	fmt.Println("│  🎉 Success: Book Added successfull!      │")
	fmt.Println("└──────────────────────────────────────┘")

}

// search book logic
func searchBooks(books *[]Books) {
	searchInput := getValidInput("Enter title to search:", 0, 0)

	for index, book := range *books {
		if strings.Contains(
			strings.ToLower(book.title),
			strings.ToLower(searchInput),
		) {
			fmt.Println("\n┌──────────────────────────────────────────────────┐")
			fmt.Printf("│ 📚 Book #%-3d                                   │\n", index+1)
			fmt.Println("├──────────────────────────────────────────────────┤")

			fmt.Printf("│ 📖 Title   : %-34s │\n", book.title)
			fmt.Printf("│ ✍ Author  : %-34s │\n", book.author)
			fmt.Printf("│ 📅 Year    : %-34s │\n", book.year)

			status := "AVAILABLE"

			if book.IsAvailable {
				status = "Borrowed"
			}

			fmt.Printf("│ 📦 Status  : %-34s │\n", status)

			fmt.Println("└──────────────────────────────────────────────────┘")
		}
	}

}

// Choose: 3
// Enter book title to borrow: The Go Programming Language
// Book is already borrowed!
func bookTitleToBorrow(books *[]Books) {

	input := getValidInput(
		"Enter book title to borrow: ",
		0,
		0,
	)

	for index, book := range *books {

		if strings.EqualFold(book.title, input) {

			if !book.IsAvailable {

				fmt.Println("\n┌──────────────────────────────────────┐")
				fmt.Println("│ ❌ Book Already Borrowed             │")
				fmt.Println("├──────────────────────────────────────┤")
				fmt.Printf("│ 📚 %-35s │\n", book.title)
				fmt.Println("└──────────────────────────────────────┘")

				return
			}

			(*books)[index].IsAvailable = false

			fmt.Println("\n┌──────────────────────────────────────┐")
			fmt.Println("│ ✅ Book Borrowed Successfully        │")
			fmt.Println("├──────────────────────────────────────┤")
			fmt.Printf("│ 📚 %-35s │\n", book.title)
			fmt.Println("└──────────────────────────────────────┘")

			return
		}
	}

	fmt.Println("\n┌──────────────────────────────────────┐")
	fmt.Println("│ ❌ No Matching Book Found            │")
	fmt.Println("└──────────────────────────────────────┘")
}

// Choose: 4
// Enter book title to return: The Go Programming Language
// Book returned successfully!
func bookTitleToReturn(books *[]Books) {

	input := getValidInput(
		"Enter book title to return: ",
		0,
		0,
	)

	for index, book := range *books {

		if strings.EqualFold(book.title, input) {

			// already available → cannot return
			if book.IsAvailable {

				fmt.Println("\n┌──────────────────────────────────────┐")
				fmt.Println("│ ❌ Book is already available        │")
				fmt.Println("├──────────────────────────────────────┤")
				fmt.Printf("│ 📚 %-35s │\n", book.title)
				fmt.Println("└──────────────────────────────────────┘")

				return
			}

			// return book
			(*books)[index].IsAvailable = true

			fmt.Println("\n┌──────────────────────────────────────┐")
			fmt.Println("│ ✅ Book returned successfully!       │")
			fmt.Println("├──────────────────────────────────────┤")
			fmt.Printf("│ 📚 %-35s │\n", book.title)
			fmt.Println("└──────────────────────────────────────┘")

			return
		}
	}

	fmt.Println("\n┌──────────────────────────────────────┐")
	fmt.Println("│ ❌ No Matching Book Found            │")
	fmt.Println("└──────────────────────────────────────┘")
}

func viewAllBooks(books *[]Books) {

	if len(*books) == 0 {
		fmt.Println("\n┌──────────────────────────────────────┐")
		fmt.Println("│ ❌ No Books Available                │")
		fmt.Println("└──────────────────────────────────────┘")
		return
	}

	for index, book := range *books {

		fmt.Println("\n┌──────────────────────────────────────────────────┐")
		fmt.Printf("│ 📚 Book #%-3d                                   │\n", index+1)
		fmt.Println("├──────────────────────────────────────────────────┤")

		fmt.Printf("│ 📖 Title   : %-34s │\n", book.title)
		fmt.Printf("│ ✍ Author  : %-34s │\n", book.author)
		fmt.Printf("│ 📅 Year    : %-34s │\n", book.year)

		status := "AVAILABLE"
		if !book.IsAvailable {
			status = "BORROWED"
		}

		fmt.Printf("│ 📦 Status  : %-34s │\n", status)

		fmt.Println("└──────────────────────────────────────────────────┘")
	}
}
