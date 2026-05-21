package main

import (
	"fmt"
)

/*
===========================================
  Go 100 Challenge
  Problem: #001
  Level: 🟢 Small
===========================================

Problem: Smart Temperature Converter

Topic:
- variables
- functions
- fmt.Scan (user input)
- if/else
- basic math

Industry Use:
IoT Sensor Systems / Weather Apps / Scientific Tools

Rules:
- User input নেবে: temperature value + unit (C, F, or K)
- বাকি 2 unit এ convert করে print করবে
- Invalid unit দিলে  => "Invalid unit! Use C, F, or K"
- Negative Kelvin দিলে => "Kelvin cannot be negative"

Formulas:
- C to F => (C * 9/5) + 32
- C to K => C + 273.15
- F to C => (F - 32) * 5/9
- K to C => K - 273.15

Example Run 1:
  Enter temperature: 100
  Enter unit (C/F/K): C
  Output: 100.00°C = 212.00°F = 373.15K

Example Run 2:
  Enter temperature: -5
  Enter unit (C/F/K): K
  Output: Kelvin cannot be negative

Example Run 3:
  Enter temperature: 100
  Enter unit (C/F/K): X
  Output: Invalid unit! Use C, F, or K
===========================================
*/

func convertTemperature(value float64, unit string) {
	// Rule 1: Validate negative Kelvin before processing
	if unit == "K" && value < 0 {
		fmt.Println("Kelvin cannot be negative")
		return
	}

	// Rule 2: Convert to all target units based on the input unit
	if unit == "C" {
		// Calculate missing units: F and K
		fahrenheit := (value * 9 / 5) + 32
		kelvin := value + 273.15
		fmt.Printf("%.2f°C = %.2f°F = %.2fK\n", value, fahrenheit, kelvin)

	} else if unit == "F" {
		// Calculate missing units: C and K
		celsius := (value - 32) * 5 / 9
		kelvin := celsius + 273.15
		fmt.Printf("%.2f°F = %.2f°C = %.2fK\n", value, celsius, kelvin)

	} else if unit == "K" {
		// Calculate missing units: C and F
		celsius := value - 273.15
		fahrenheit := (celsius * 9 / 5) + 32
		fmt.Printf("%.2fK = %.2f°C = %.2f°F\n", value, celsius, fahrenheit)

	} else {
		// Rule 3: Error handle invalid entries exactly as requested
		fmt.Println("Invalid unit! Use C, F, or K")
	}
}

func main() {
	var temperature float64
	var unit string

	fmt.Println("Enter temperature:")
	fmt.Scan(&temperature)

	fmt.Println("Enter unit (C/F/K):")
	fmt.Scan(&unit)

	convertTemperature(temperature, unit)
}
