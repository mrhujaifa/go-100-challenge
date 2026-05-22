package main

import (
	"fmt"
	"strings"
)

/*
===========================================
  Go 100 Challenge
  Problem: #005
  Level: 🟢 Small
===========================================

Problem: Student Grade Calculator

Topic:
- struct
- methods on struct
- switch statement
- average calculation

Industry Use:
School Management System / LMS /
Education Platforms / Report Card Generator

Rules:
- Student struct বানাবে (Name, Marks []float64)
- Struct এ method বানাবে => Average() float64
- Struct এ method বানাবে => Grade() string
- Grade logic:
    90-100 => "A+"
    80-89  => "A"
    70-79  => "B"
    60-69  => "C"
    50-59  => "D"
    0-49   => "F"
- Result print করবে cleanly

Example Output:
  ================================
  Student:  Rakib
  Marks:    85, 90, 78, 92, 88
  Average:  86.60
  Grade:    A
  ================================
===========================================
*/

type Student struct {
	Name  string
	Marks []float64
}

func (s Student) Average() float64 {
	if len(s.Marks) == 0 {
		return 0
	}

	var total float64
	for _, mark := range s.Marks {
		total += mark
	}

	return total / float64(len(s.Marks))
}

func (s Student) Grade() string {
	avg := s.Average()

	switch {
	case avg >= 90:
		return "A+"
	case avg >= 80:
		return "A"
	case avg >= 70:
		return "B"
	case avg >= 60:
		return "C"
	case avg >= 50:
		return "D"
	default:
		return "F"
	}
}

func printResult(s Student) {
	// Marks গুলো comma দিয়ে join করো
	parts := make([]string, len(s.Marks))
	for i, m := range s.Marks {
		parts[i] = fmt.Sprintf("%.0f", m)
	}
	marksStr := strings.Join(parts, ", ")

	fmt.Println("  ================================")
	fmt.Printf("  Student:  %s\n", s.Name)
	fmt.Printf("  Marks:    %s\n", marksStr)
	fmt.Printf("  Average:  %.2f\n", s.Average())
	fmt.Printf("  Grade:    %s\n", s.Grade())
}

func main() {
	students := []Student{
		{Name: "Rakib", Marks: []float64{85, 90, 78, 92, 88}},
		{Name: "Sumaiya", Marks: []float64{45, 50, 38, 60, 42}},
		{Name: "Tanvir", Marks: []float64{70, 75, 80, 65, 72}},
	}

	for _, s := range students {
		printResult(s)
	}
	fmt.Println("  ================================")
}
