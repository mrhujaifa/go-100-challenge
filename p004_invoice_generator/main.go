package main

import "fmt"

/*
===========================================
  Go 100 Challenge
  Problem: #004
  Level: 🟢 Small
===========================================

Problem: Invoice Generator

Topic:
- struct
- slice of structs
- for range loop
- fmt formatting

Industry Use:
E-commerce / POS Systems / Billing Software /
Accounting Tools

Rules:
- একটা Product struct বানাবে (Name, Price, Quantity)
- কমপক্ষে 3টা product hardcode করবে
- প্রতিটা product এর subtotal = Price * Quantity
- সবশেষে Total print করবে
- 15% VAT যোগ করবে total এর উপর
- Clean formatted output দেবে

Example Output:
  ================================
         INVOICE
  ================================
  Product       Qty   Price   Sub
  --------------------------------
  Apple          3    20.00   60.00
  Banana         5    10.00   50.00
  Mango          2    50.00  100.00
  --------------------------------
  Subtotal:             210.00
  VAT (15%):             31.50
  ================================
  TOTAL:                241.50
  ================================
===========================================
*/

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func generateInvoice(products []Product) {
	fmt.Println("========================================")
	fmt.Println("                INVOICE                 ")
	fmt.Println("========================================")

	fmt.Printf("%-14s %-5s %-9s %s\n", "Product", "Qty", "Price", "Sub")
	fmt.Println("----------------------------------------")

	var subtotal float64 = 0

	// all products looping
	for _, p := range products {
		sub := float64(p.Quantity) * p.Price
		subtotal += sub

		fmt.Printf("%-14s %-5d %-9.2f %6.2f\n", p.Name, p.Quantity, p.Price, sub)
	}

	vat := subtotal * 0.15
	total := subtotal + vat

	fmt.Println("----------------------------------------")
	fmt.Printf("%-30s %6.2f\n", "Subtotal:", subtotal)
	fmt.Printf("%-30s %6.2f\n", "VAT (15%):", vat)
	fmt.Println("========================================")
	fmt.Printf("%-30s %6.2f\n", "TOTAL:", total)
	fmt.Println("========================================")
}

func main() {
	products := []Product{
		{
			Name:     "Apple",
			Price:    30,
			Quantity: 5,
		},
		{
			Name:     "Apple",
			Price:    30,
			Quantity: 5,
		},
		{
			Name:     "Apple",
			Price:    30,
			Quantity: 5,
		},
	}

	generateInvoice(products)
}
