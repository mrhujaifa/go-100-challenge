package main

import (
	"fmt"
	"strings"
	"time"
)

/*
===========================================
  Go 100 Challenge
  Problem: #017
  Level: 🟡 Medium
===========================================

Problem: Employee Management System

Topic:
- struct
- map
- functions
- interface
- error handling

Industry Use:
HR Systems / Payroll Management /
Company Portals / Employee Tracking

Rules (English):
- Manage employees with salary and department
- Each employee has: ID, Name, Department, Salary, JoiningDate
- User can:
    * Add employee
    * View all employees
    * Search by department
    * Give salary raise (%)
    * Fire employee (remove)
    * View highest paid employee
    * Exit
- Duplicate ID          => "Employee ID already exists!"
- Employee not found    => "Employee not found!"
- Raise must be 1-50%   => "Raise must be between 1% and 50%!"
- Empty department      => "No employees in this department!"

Rules (বাংলায়):
- Salary আর department সহ employees manage করবে
- প্রতিটা employee এ থাকবে: ID, Name, Department, Salary, JoiningDate
- User করতে পারবে:
    * Employee add করা
    * সব employees দেখা
    * Department দিয়ে search করা
    * Salary raise দেওয়া (%)
    * Employee fire করা (remove)
    * সবচেয়ে বেশি বেতনের employee দেখা
    * Exit করা
- Duplicate ID           => "Employee ID already exists!"
- Employee না পেলে      => "Employee not found!"
- Raise 1-50% এর বাইরে => "Raise must be between 1% and 50%!"
- Department খালি       => "No employees in this department!"

Example Run:
  === Employee Management ===
  1. Add Employee
  2. View All
  3. Search By Department
  4. Give Raise
  5. Fire Employee
  6. Highest Paid
  7. Exit

  Choose: 1
  ID: EMP001
  Name: Rakib
  Department: Engineering
  Salary: 50000
  ✅ Employee added!

  Choose: 4
  Employee ID: EMP001
  Raise %: 20
  ✅ Salary updated!
  Rakib: 50000.00 → 60000.00

  Choose: 6
  👑 Highest Paid Employee:
  EMP001 | Rakib | Engineering | 60000.00
===========================================
*/

const (
	MinSalary = 20000.0
	MaxSalary = 50000.0
	MinRaise  = 1.0
	MaxRaise  = 50.0
)

func main() {
	employees := &map[string]Employee{}
	MainMenu(employees)
}

type Employee struct {
	ID          string
	Name        string
	Department  string
	Salary      float64
	JoiningDate string
}

func MainMenu(employees *map[string]Employee) {
	var choice int

	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════╗")
		fmt.Println("║         👥 EMPLOYEE MANAGEMENT SYSTEM        ║")
		fmt.Println("╠══════════════════════════════════════════════╣")
		fmt.Println("║                                              ║")
		fmt.Println("║   [1] ➕ Add Employee                        ║")
		fmt.Println("║   [2] 👀 View All Employees                  ║")
		fmt.Println("║   [3] 🔍 Search By Department                ║")
		fmt.Println("║   [4] 💰 Give Salary Raise                   ║")
		fmt.Println("║   [5] 🔥 Fire Employee                       ║")
		fmt.Println("║   [6] 👑 Highest Paid Employee               ║")
		fmt.Println("║                                              ║")
		fmt.Println("║   [0] 🚪 Exit                                ║")
		fmt.Println("║                                              ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Print("\n👉 Choose an option: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			if err := addEmployee(employees); err != nil {
				fmt.Println("❌", err)
			}
		case 2:
			if err := viewAllEmployees(employees); err != nil {
				fmt.Println("❌", err)
			}
		case 3:
			if err := searchByDepartment(employees); err != nil {
				fmt.Println("❌", err)
			}
		case 4:
			if err := giveSalaryRaise(employees); err != nil {
				fmt.Println("❌", err)
			}
		case 5:
			if err := fireEmployee(employees); err != nil {
				fmt.Println("❌", err)
			}
		case 6:
			if err := highestPaid(employees); err != nil {
				fmt.Println("❌", err)
			}
		case 0:
			fmt.Println("\n👋 Goodbye!")
			return
		default:
			fmt.Println("❌ Invalid option!")
		}
	}
}

// readEmployeeInput gathers employee data from the console
// and returns the collected values.
func readEmployeeInput() (string, string, string, float64, error) {
	var id, name, department string
	var salary float64

	fmt.Print("ID: ")
	fmt.Scan(&id)

	fmt.Print("Name: ")
	fmt.Scan(&name)

	fmt.Print("Department: ")
	fmt.Scan(&department)

	fmt.Print("Salary: ")
	fmt.Scan(&salary)

	return id, name, department, salary, nil
}

// validateEmployee validates employee details before processing.
func validateEmployee(id, name, department string, salary float64) error {
	if id == "" {
		return fmt.Errorf("ID cannot be empty")
	}

	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if department == "" {
		return fmt.Errorf("department cannot be empty")
	}

	if salary < MinSalary {
		return fmt.Errorf("salary must be at least %.2f", MinSalary)
	}

	if salary > MaxSalary {
		return fmt.Errorf("salary cannot exceed %.2f", MaxSalary)
	}

	return nil
}

func saveEmployee(employees *map[string]Employee, id, name, department string, salary float64) error {

	if _, exists := (*employees)[id]; exists {
		return fmt.Errorf("employee ID %s already exists", id)
	}

	(*employees)[id] = Employee{
		ID:          id,
		Name:        name,
		Department:  department,
		Salary:      salary,
		JoiningDate: time.Now().Format("2006-01-02"),
	}

	return nil
}

func addEmployee(employees *map[string]Employee) error {

	id, name, department, salary, err := readEmployeeInput()
	if err != nil {
		return fmt.Errorf("input error: %w", err)
	}

	if err := validateEmployee(id, name, department, salary); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := saveEmployee(employees, id, name, department, salary); err != nil {
		return fmt.Errorf("save error: %w", err)
	}

	fmt.Printf("✅ Employee %s added successfully!\n", name)
	return nil
}

func viewAllEmployees(employees *map[string]Employee) error {
	if len(*employees) == 0 {
		return fmt.Errorf("no employees found")
	}

	printEmployeeHeader()

	for _, e := range *employees {
		printEmployeeRow(e)
	}

	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	fmt.Printf("  Total Employees: %d\n", len(*employees))

	return nil
}

func printEmployeeHeader() {
	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│                    👥 ALL EMPLOYEES                         │")
	fmt.Println("├──────────┬────────────────┬────────────────┬───────────────┤")
	fmt.Printf("│ %-8s │ %-14s │ %-14s │ %-13s │\n",
		"ID", "Name", "Department", "Salary")
	fmt.Println("├──────────┼────────────────┼────────────────┼───────────────┤")
}

func printEmployeeRow(e Employee) {
	fmt.Printf("│ %-8s │ %-14s │ %-14s │ %13.2f │\n",
		e.ID,
		e.Name,
		e.Department,
		e.Salary,
	)
}

func searchByDepartment(employees *map[string]Employee) error {
	// collect department input from user
	var department string
	fmt.Print("Search Department: ")
	fmt.Scan(&department)

	// validate — department cannot be empty
	if department == "" {
		return fmt.Errorf("department cannot be empty")
	}

	// search employees by department
	found := false
	printEmployeeHeader()

	for _, e := range *employees {
		// case-insensitive match
		if strings.EqualFold(e.Department, department) {
			printEmployeeRow(e)
			found = true
		}
	}

	fmt.Println("└─────────────────────────────────────────────────────────────┘")

	// return error if no match found
	if !found {
		return fmt.Errorf("no employees found in department: %s", department)
	}

	return nil
}

// giveSalaryRaise increases the salary of an employee
// by the given percentage. Raise must be between MinRaise and MaxRaise.
// Returns an error if employee not found or raise is invalid.
func giveSalaryRaise(employees *map[string]Employee) error {

	// collect employee ID from user
	var empId string
	fmt.Print("Employee ID: ")
	fmt.Scan(&empId)

	// check employee exists
	employee, exists := (*employees)[empId]
	if !exists {
		return fmt.Errorf("employee %s not found", empId)
	}

	// collect raise percentage from user
	var raise float64
	fmt.Print("Raise %%: ")
	fmt.Scan(&raise)

	// validate raise range
	if raise < MinRaise || raise > MaxRaise {
		return fmt.Errorf("raise must be between %.0f%% and %.0f%%", MinRaise, MaxRaise)
	}

	// calculate new salary
	oldSalary := employee.Salary
	newSalary := oldSalary * (1 + raise/100)

	// update salary in map
	employee.Salary = newSalary
	(*employees)[empId] = employee

	// print result
	fmt.Printf("✅ Salary Updated!\n")
	fmt.Printf("%s: %.2f → %.2f\n", employee.Name, oldSalary, newSalary)

	return nil
}

// fireEmployee removes an employee from the system by ID.
// Returns an error if the employee is not found.
func fireEmployee(employees *map[string]Employee) error {

	// collect employee ID from user
	var empId string
	fmt.Print("Employee ID to fire: ")
	fmt.Scan(&empId)

	// check employee exists
	employee, exists := (*employees)[empId]
	if !exists {
		return fmt.Errorf("employee %s not found", empId)
	}

	// confirm before delete
	var confirm string
	fmt.Printf("⚠️  Are you sure you want to fire %s? (yes/no): ", employee.Name)
	fmt.Scan(&confirm)

	if !strings.EqualFold(confirm, "yes") {
		fmt.Println("❌ Cancelled!")
		return nil
	}

	// remove employee from map
	delete(*employees, empId)

	fmt.Printf("✅ Employee %s has been fired!\n", employee.Name)
	return nil
}

// highestPaid finds and displays the employee
// with the highest salary in the system.
// Returns an error if no employees exist.
func highestPaid(employees *map[string]Employee) error {

	// check employees exist
	if len(*employees) == 0 {
		return fmt.Errorf("no employees found")
	}

	// find highest paid employee
	var topEmployee Employee

	for _, e := range *employees {
		if e.Salary > topEmployee.Salary {
			topEmployee = e
		}
	}

	// print result
	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│                  👑 HIGHEST PAID EMPLOYEE                   │")
	fmt.Println("├──────────┬────────────────┬────────────────┬───────────────┤")
	fmt.Printf("│ %-8s │ %-14s │ %-14s │ %-13s │\n",
		"ID", "Name", "Department", "Salary")
	fmt.Println("├──────────┼────────────────┼────────────────┼───────────────┤")
	printEmployeeRow(topEmployee)
	fmt.Println("└─────────────────────────────────────────────────────────────┘")

	return nil
}
