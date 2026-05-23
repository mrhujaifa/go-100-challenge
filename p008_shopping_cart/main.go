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
  Problem: #008
  Level: 🟢 Small
===========================================

Problem: Shopping Cart System

Topic:
- struct
- slice of structs
- map
- functions
- for loop

Industry Use:
E-commerce / Retail Apps /
POS Systems / Online Shopping

Rules:
- Create a cart where user can add items
- Each item has: Name, Price, Quantity
- User can:
    * Add item to cart
    * Remove item by name
    * View cart with all items
    * Checkout — shows total price
- Empty cart checkout => "Cart is empty!"
- Remove item not in cart => "Item not found!"

বাংলায়:
- একটা shopping cart বানাবে
- প্রতিটা item এ থাকবে: Name, Price, Quantity
- User করতে পারবে:
    * Cart এ item add করা
    * নাম দিয়ে item remove করা
    * Cart এর সব item দেখা
    * Checkout — মোট দাম দেখাবে
- খালি cart এ checkout => "Cart is empty!"
- নেই এমন item remove => "Item not found!"

Example Run:
  === Shopping Cart ===
  1. Add Item
  2. Remove Item
  3. View Cart
  4. Checkout
  5. Exit

  Choose: 1
  Item name: Apple
  Price: 20
  Quantity: 3
  Item added!

  Choose: 3
  Apple      x3   60.00

  Choose: 4
  ========================
  Total: 60.00
  ========================
===========================================
*/

type CartItems struct {
	Name     string
	Price    float64
	Quantity int
}

var reader = bufio.NewReader(os.Stdin)

func readInput() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	MyCart := []CartItems{}
	var choice int
	for {
		fmt.Println("=== Shopping Cart ===")
		fmt.Println("1. Add Item")
		fmt.Println("2. Remove Item")
		fmt.Println("3. View Cart")
		fmt.Println("4. Checkout")
		fmt.Println("5. Exit")
		fmt.Print("Choose: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			MyCart = addToCart(MyCart)
		case 2:
			MyCart = RemoveFromCart(MyCart)
		case 3:
			MyCart = viewCartItems(MyCart)
		case 4:
			MyCart = checkout(MyCart)
		case 5:
			return
		}
	}

}

func addToCart(items []CartItems) []CartItems {

	fmt.Print("Item name: ")
	nameInput := readInput()

	fmt.Print("Price: ")
	priceInput := readInput()

	var price float64
	fmt.Sscan(priceInput, &price)

	fmt.Print("Quantity: ")
	quantInput := readInput()

	var quantity int
	fmt.Sscan(quantInput, &quantity)

	item := CartItems{
		Name:     nameInput,
		Price:    price,
		Quantity: quantity,
	}

	newItems := append(items, item)

	fmt.Println("Item added!", newItems)
	return newItems
}

func viewCartItems(items []CartItems) []CartItems {

	if len(items) == 0 {
		fmt.Println("no items here!")
	}
	for i, item := range items {
		fmt.Println("CartNum:", i, item)
	}

	return items
}

func RemoveFromCart(items []CartItems) []CartItems {

	fmt.Print("Enter your item name for remove item: ")
	nameInput := readInput()

	if len(nameInput) == 0 {
		fmt.Println("Invalid input")
		return items
	}

	var updatedCart []CartItems
	found := false

	for _, item := range items {

		if item.Name == nameInput {
			found = true
			continue
		}

		updatedCart = append(updatedCart, item)
	}

	if !found {
		fmt.Println("Item not found!")
		return items
	}

	fmt.Println("Item removed successfully!")
	return updatedCart
}

func checkout(items []CartItems) []CartItems {

	if len(items) == 0 {
		fmt.Println("No items in cart")
		return items
	}

	var totalCheckout float64

	for _, item := range items {

		total := float64(item.Quantity) * item.Price

		totalCheckout += total
	}

	fmt.Println("Total Checkout:", totalCheckout)

	return items
}
