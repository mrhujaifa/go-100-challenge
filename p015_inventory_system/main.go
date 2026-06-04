package main

import (
	"fmt"
	"strings"
)

/*
===========================================
  Go 100 Challenge
  Problem: #015
  Level: 🟡 Medium
===========================================

Problem: Inventory Management System

Topic:
- struct
- map
- functions
- stock management
- error handling

Industry Use:
Warehouse Systems / E-commerce /
Retail Management / Supply Chain

Rules (English):
- Manage product inventory with stock levels
- Each product has: ID, Name, Price, Stock, Category
- User can:
    * Add new product
    * Update stock (add or remove)
    * View all products
    * Search by category
    * Low stock alert (stock < 10)
    * Exit
- Product already exists  => "Product already exists!"
- Product not found       => "Product not found!"
- Remove more than stock  => "Insufficient stock!"
- Low stock alert         => show all products with stock < 10

Rules (বাংলায়):
- Product inventory stock সহ manage করবে
- প্রতিটা product এ থাকবে: ID, Name, Price, Stock, Category
- User করতে পারবে:
    * নতুন product add করা
    * Stock update করা (add বা remove)
    * সব products দেখা
    * Category দিয়ে search করা
    * Low stock alert দেখা (stock < 10)
    * Exit করা
- Product already exists  => "Product already exists!"
- Product না পেলে         => "Product not found!"
- Stock এর বেশি remove   => "Insufficient stock!"
- Low stock alert         => stock < 10 এমন সব products দেখাবে

Example Run:
  === Inventory System ===
  1. Add Product
  2. Update Stock
  3. View All Products
  4. Search By Category
  5. Low Stock Alert
  6. Exit

  Choose: 1
  Name: Rice
  Price: 50.00
  Stock: 100
  Category: Food
  ✅ Product added! ID: PRD001

  Choose: 2
  Product ID: PRD001
  1. Add Stock
  2. Remove Stock
  Choose: 2
  Quantity: 95
  ✅ Stock updated! Rice - Remaining: 5

  Choose: 5
  ⚠️  Low Stock Alert!
  PRD001  Rice   50.00   5    Food
===========================================
*/

type Product struct {
	ID       string
	name     string
	price    float64
	stock    int
	category string
}

func main() {
	products := &map[string]Product{}

	MainMenu(products)
}

// main menu CLI
func MainMenu(products *map[string]Product) {

	var choice int

	for {
		fmt.Println()
		fmt.Println("======================================================")
		fmt.Println("               === Inventory System ===")
		fmt.Println("======================================================")
		fmt.Println()
		fmt.Println("  [1] Add Product")
		fmt.Println("  [2] Update Stock")
		fmt.Println("  [3] View All Products")
		fmt.Println("  [4] Search By Category")
		fmt.Println("  [5] Low Stock Alert")
		fmt.Println()
		fmt.Println("  [0] Exit")
		fmt.Println()
		fmt.Println("------------------------------------------------------")
		fmt.Print(" Select Option > ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			createProduct(products)
		case 2:
			updateStock(products)
		case 3:
			viewAllProducts(products)
		case 4:
			searchByCategory(products)
		case 5:
			lowStockAlert(products)
		case 0:
			return
		default:
			fmt.Println("Invlaid Input!")
		}

	}
}

// create product
func createProduct(products *map[string]Product) {

	var (
		name     string
		category string
		price    float64
		stock    int
	)

	// input
	fmt.Print("Name: ")
	fmt.Scan(&name)

	fmt.Print("Price: ")
	fmt.Scan(&price)

	fmt.Print("Stock: ")
	fmt.Scan(&stock)

	fmt.Print("Category: ")
	fmt.Scan(&category)

	// ID generate
	id := fmt.Sprintf("PRD%03d", len(*products)+1)

	// created new product
	newProduct := Product{
		ID:       id,
		name:     name,
		price:    price,
		stock:    stock,
		category: category,
	}

	// append
	(*products)[id] = newProduct

	fmt.Printf("✅ Product added! ID: %s\n", id)
}

// update product stock
func updateStock(products *map[string]Product) {

	var (
		ID     string
		choice int
	)

	fmt.Println("Enter your Product ID: ")
	fmt.Scan(&ID)

	fmt.Println("1. Add Stock")
	fmt.Println("2. Remove Stock")
	fmt.Scan(&choice)

	var (
		quantity int
	)

	fmt.Println("Quantity: ")
	fmt.Scan(&quantity)

	switch choice {
	case 1:

		product, exists := (*products)[ID]

		if !exists {
			fmt.Println("❌ Product not found!")
			return
		}

		product.stock = quantity

		(*products)[ID] = product

		fmt.Println("✅ Stock updated:", product.name, "- Remaining:", product.stock)

	case 2:

		product, exists := (*products)[ID]

		if !exists {
			fmt.Println("❌ Product not found!")
			return
		}

		if quantity > product.stock {
			fmt.Println("❌ Not enough stock available!")
			return
		}

		product.stock -= quantity

		(*products)[ID] = product

		fmt.Println("✅ Stock updated:", product.name, "- Remaining:", product.stock)

	}

}

// view all products
func viewAllProducts(products *map[string]Product) {
	for _, p := range *products {
		fmt.Println("----------------------")
		fmt.Println("ID:", p.ID)
		fmt.Println("productName: ", p.name)
		fmt.Println("categories: ", p.category)
		fmt.Println("Price: ", p.price)
		fmt.Println("Stock: ", p.stock)
		fmt.Println("----------------------")
	}
}

// Product search by category
func searchByCategory(products *map[string]Product) {
	if products == nil || len(*products) == 0 {
		fmt.Println("❌ No products available.")
		return
	}

	var search string

	fmt.Print("Enter category to search: ")
	fmt.Scanln(&search)

	search = strings.TrimSpace(search)

	if search == "" {
		fmt.Println("❌ Category cannot be empty.")
		return
	}

	search = strings.ToLower(search)

	found := false

	for _, p := range *products {
		category := strings.ToLower(strings.TrimSpace(p.category))

		if strings.Contains(category, search) {
			fmt.Printf("✅ ID: %s | Name: %s | Category: %s | Stock: %d\n",
				p.ID,
				p.name,
				p.category,
				p.stock,
			)
			found = true
		}
	}

	if !found {
		fmt.Printf("❌ No products found for category '%s'\n", search)
	}
}

func lowStockAlert(products *map[string]Product) {
	const LowStockThreshold = 5

	if products == nil || len(*products) == 0 {
		fmt.Println("❌ No products available.")
		return
	}

	found := false

	fmt.Println("\n⚠️ LOW STOCK ALERT")
	fmt.Println("========================================================")
	fmt.Printf("%-10s %-15s %-10s %-10s %-15s\n",
		"ID", "NAME", "PRICE", "STOCK", "CATEGORY")
	fmt.Println("--------------------------------------------------------")

	for _, p := range *products {
		if p.stock <= LowStockThreshold {
			fmt.Printf("%-10s %-15s %-10.2f %-10d %-15s\n",
				p.ID,
				p.name,
				p.price,
				p.stock,
				p.category,
			)
			found = true
		}
	}

	if !found {
		fmt.Println("✅ No low stock products found.")
	}

	fmt.Println("========================================================")
}
