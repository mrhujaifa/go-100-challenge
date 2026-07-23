package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

/*
===========================================
  Go 100 Challenge
  Problem: #020
  Level: 🟡 Medium
===========================================

Problem: Subscription Plan Manager

Topic:
- struct
- interface (first time!)
- custom type
- functions
- error handling

Industry Use:
SaaS Platforms / Streaming Services /
Membership Apps / Billing Systems

Rules (English):
- Manage user subscriptions with different plans
- Plans: Free, Basic, Pro, Enterprise
- Each plan has different limits:
    * Free       → 5 projects,  1GB storage,  no support
    * Basic      → 20 projects, 10GB storage, email support
    * Pro        → 100 projects, 50GB storage, priority support
    * Enterprise → unlimited projects, 500GB storage, 24/7 support
- Each user has: Name, Email, Plan, StartDate
- User can:
    * Register user with a plan
    * Upgrade plan
    * Downgrade plan
    * View user details + plan limits
    * View all users
    * Exit
- User not found        => "User not found!"
- Invalid plan          => "Invalid plan!"
- Already on same plan  => "User already on this plan!"
- Downgrade to Free     => "Cannot downgrade to Free plan directly!"

Rules (বাংলায়):
- Different plans সহ user subscriptions manage করবে
- Plans: Free, Basic, Pro, Enterprise
- প্রতিটা plan এর আলাদা limits:
    * Free       → 5 projects,  1GB storage,  no support
    * Basic      → 20 projects, 10GB storage, email support
    * Pro        → 100 projects, 50GB storage, priority support
    * Enterprise → unlimited projects, 500GB storage, 24/7 support
- প্রতিটা user এ থাকবে: Name, Email, Plan, StartDate
- User করতে পারবে:
    * Plan সহ user register করা
    * Plan upgrade করা
    * Plan downgrade করা
    * User details + plan limits দেখা
    * সব users দেখা
    * Exit করা
- User না পেলে          => "User not found!"
- Invalid plan           => "Invalid plan!"
- Same plan এ থাকলে    => "User already on this plan!"
- Free তে downgrade     => "Cannot downgrade to Free plan directly!"

Example Run:
  === Subscription Manager ===
  1. Register User
  2. Upgrade Plan
  3. Downgrade Plan
  4. View User
  5. View All Users
  6. Exit

  Choose: 1
  Name : Rakib
  Email: rakib@gmail.com
  Plan (Free/Basic/Pro/Enterprise): Basic
  ✅ User registered!

  Choose: 4
  Email: rakib@gmail.com
  ┌──────────────────────────────────────┐
  │ Name    : Rakib                      │
  │ Email   : rakib@gmail.com            │
  │ Plan    : Basic                      │
  │ Start   : 2024-01-15                 │
  ├──────────────────────────────────────┤
  │ PLAN LIMITS                          │
  │ Projects : 20                        │
  │ Storage  : 10GB                      │
  │ Support  : Email                     │
  └──────────────────────────────────────┘

  Choose: 2
  Email: rakib@gmail.com
  New Plan: Pro
  ✅ Upgraded to Pro!
===========================================
*/

// Planner defines the interface for subscription plans.
// Every plan must implement these methods.
type Planner interface {
	PlanName() string
	MaxProjects() int
	StorageGB() int
	SupportType() string
}

type User struct {
	Name      string
	Email     string
	Plan      string
	StartDate string
}

// constants
const (
	PlanFree       = "Free"
	PlanBasic      = "Basic"
	PlanPro        = "Pro"
	PlanEnterprise = "Enterprise"
)

// Custom Errors
var (
	ErrUserNotFound     = errors.New("User not found!")
	ErrInvalidPlan      = errors.New("Invalid plan!")
	ErrAlreadyOnPlan    = errors.New("User already on this plan!")
	ErrDowngradeToFree  = errors.New("Cannot downgrade to Free plan directly!")
	ErrEmptyFields      = errors.New("Name and Email cannot be empty!")
	ErrUserAlreadyExist = errors.New("User with this email already exists!")
)

// Plan Implementation
type FreePlan struct{}
type BasicPlan struct{}
type ProPlan struct{}
type EnterprisePlan struct{}

func (f FreePlan) PlanName() string    { return PlanFree }
func (f FreePlan) MaxProjects() int    { return 5 }
func (f FreePlan) StorageGB() int      { return 1 }
func (f FreePlan) SupportType() string { return "no support" }

func (b BasicPlan) PlanName() string    { return PlanBasic }
func (b BasicPlan) MaxProjects() int    { return 20 }
func (b BasicPlan) StorageGB() int      { return 10 }
func (b BasicPlan) SupportType() string { return "email support" }

func (p ProPlan) PlanName() string    { return PlanPro }
func (p ProPlan) MaxProjects() int    { return 100 }
func (p ProPlan) StorageGB() int      { return 50 }
func (p ProPlan) SupportType() string { return "priority support" }

func (e EnterprisePlan) PlanName() string    { return PlanEnterprise }
func (e EnterprisePlan) MaxProjects() int    { return -1 } // -1 indicates unlimited
func (e EnterprisePlan) StorageGB() int      { return 500 }
func (e EnterprisePlan) SupportType() string { return "24/7 support" }

type UserRepository struct {
	mu    sync.RWMutex //TODO : learning need
	users map[string]*User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*User),
	}
}
func GetPlan(planName string) (Planner, error) {
	switch strings.ToLower(planName) {
	case "free":
		return FreePlan{}, nil
	case "basic":
		return BasicPlan{}, nil
	case "pro":
		return ProPlan{}, nil
	case "enterprise":
		return EnterprisePlan{}, nil
	default:
		return nil, ErrInvalidPlan
	}
}

func (r *UserRepository) UserRegister(name, email, planName string) error {

	// input fields validation
	if strings.TrimSpace(name) == "" || strings.TrimSpace(email) == "" {
		return ErrEmptyFields
	}

	// plan checking validation
	planNameTitle := strings.Title(strings.ToLower(planName))

	if _, err := GetPlan(planNameTitle); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// Email existing validation
	if _, exist := r.users[email]; exist {
		return ErrUserAlreadyExist
	}

	// insert data
	r.users[email] = &User{
		Name:      name,
		Email:     email,
		Plan:      planNameTitle,
		StartDate: time.Now().Format("2006-01-02"),
	}

	return nil

}

func (r *UserRepository) Upgrade(email, planName string) error {

	r.mu.Lock()

	defer r.mu.Unlock()

	// input fields validation
	if strings.TrimSpace(email) == "" || strings.TrimSpace(planName) == "" {
		return ErrEmptyFields
	}

	newPlanTitle := strings.Title(strings.ToLower(planName))
	if _, err := GetPlan(newPlanTitle); err != nil {
		return ErrInvalidPlan
	}

	for _, user := range r.users {
		if strings.EqualFold(user.Email, email) {
			user.Plan = planName
		}
	}
	return nil
}

func printUserDetails(user *User) {
	plan, _ := GetPlan(user.Plan)

	projectsStr := fmt.Sprintf("%d", plan.MaxProjects())
	if plan.MaxProjects() == -1 {
		projectsStr = "Unlimited"
	}

	fmt.Printf("\n  ┌──────────────────────────────────────┐\n")
	fmt.Printf("  │ Name     : %-25s │\n", user.Name)
	fmt.Printf("  │ Email    : %-25s │\n", user.Email)
	fmt.Printf("  │ Plan     : %-25s │\n", user.Plan)
	fmt.Printf("  │ Start    : %-25s │\n", user.StartDate)
	fmt.Printf("  ├──────────────────────────────────────┤\n")
	fmt.Printf("  │ PLAN LIMITS                          │\n")
	fmt.Printf("  │ Projects : %-25s │\n", projectsStr)
	fmt.Printf("  │ Storage  : %-23dGB │\n", plan.StorageGB())
	fmt.Printf("  │ Support  : %-25s │\n", plan.SupportType())
	fmt.Printf("  └──────────────────────────────────────┘\n\n")
}

func (r *UserRepository) GetUser(email string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[email]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}

// input func
var scanner = bufio.NewScanner(os.Stdin)

func getInput(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func main() {

	repo := NewUserRepository()

	for {
		fmt.Println("\n=== Subscription Manager ===")
		fmt.Println("1. Register User")
		fmt.Println("2. Upgrade Plan")
		fmt.Println("3. Downgrade Plan")
		fmt.Println("4. View User")
		fmt.Println("5. View All Users")
		fmt.Println("6. Exit")

		choice := getInput("\nChoose: ")

		switch choice {
		case "1":
			name := getInput("Name : ")
			email := getInput("Email: ")
			plan := getInput("Plan (Free/Basic/Pro/Enterprise): ")

			if err := repo.UserRegister(name, email, plan); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("User registered succesfully!")
			}

		case "2":
			email := getInput("Enter Email:")
			newPlanName := getInput("Enter Upgrade Planname:")

			if err := repo.Upgrade(email, newPlanName); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Plan upgrade succesfully!")
			}

		case "3":
			email := getInput("Email: ")
			user, err := repo.GetUser(email)
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			} else {
				printUserDetails(user)
			}

		}

		fmt.Println(repo)

	}

}
