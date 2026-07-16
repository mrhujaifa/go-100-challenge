package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
===========================================
  Go 100 Challenge
  Problem: #018
  Level: 🟡 Medium
===========================================

Problem: Recipe Manager CLI

Topic:
- struct
- slice
- map
- functions
- error handling

Industry Use:
Food Apps / Restaurant Systems /
Meal Planning / Nutrition Tracking

Rules (English):
- Manage cooking recipes with ingredients
- Each recipe has: Name, Category, CookTime, Ingredients, Steps
- Categories: Breakfast, Lunch, Dinner, Snack, Dessert
- User can:
    * Add recipe
    * View all recipes
    * Search by category
    * View recipe details (ingredients + steps)
    * Delete recipe
    * Exit
- Duplicate recipe name   => "Recipe already exists!"
- Recipe not found        => "Recipe not found!"
- Invalid category        => "Invalid category!"
- Empty ingredients       => "Recipe must have at least 1 ingredient!"

Rules (বাংলায়):
- Cooking recipes ingredients সহ manage করবে
- প্রতিটা recipe এ থাকবে: Name, Category, CookTime, Ingredients, Steps
- Categories: Breakfast, Lunch, Dinner, Snack, Dessert
- User করতে পারবে:
    * Recipe add করা
    * সব recipes দেখা
    * Category দিয়ে search করা
    * Recipe details দেখা (ingredients + steps)
    * Recipe delete করা
    * Exit করা
- Duplicate recipe name  => "Recipe already exists!"
- Recipe না পেলে        => "Recipe not found!"
- Invalid category       => "Invalid category!"
- Empty ingredients      => "Recipe must have at least 1 ingredient!"

Example Run:
  === Recipe Manager ===
  1. Add Recipe
  2. View All
  3. Search By Category
  4. View Details
  5. Delete Recipe
  6. Exit

  Choose: 1
  Name     : Pasta
  Category : Dinner
  CookTime : 30
  Ingredients (comma separated): pasta, sauce, cheese
  Steps (comma separated): boil pasta, add sauce, add cheese
  ✅ Recipe added!

  Choose: 4
  Recipe Name: Pasta
  ┌─────────────────────────────────┐
  │ Name     : Pasta                │
  │ Category : Dinner               │
  │ CookTime : 30 mins              │
  ├─────────────────────────────────┤
  │ Ingredients:                    │
  │  - pasta                        │
  │  - sauce                        │
  │  - cheese                       │
  ├─────────────────────────────────┤
  │ Steps:                          │
  │  1. boil pasta                  │
  │  2. add sauce                   │
  │  3. add cheese                  │
  └─────────────────────────────────┘
===========================================
*/

type Recipe struct {
	Name        string
	Category    string
	CookTime    int
	Ingredients []string
	Steps       []string
}

type RecipeRepository struct {
	data map[string]Recipe
}

// Constructor
func NewRecipeRepository() *RecipeRepository {
	return &RecipeRepository{
		data: make(map[string]Recipe),
	}
}

// Methods
func (r *RecipeRepository) Add(recipe Recipe) error {
	if _, exists := r.data[recipe.Name]; exists {
		return errors.New("recipe already exists")
	}

	if recipe.CookTime == 0 {
		return errors.New("CookTime must be require!")
	}

	if len(recipe.Ingredients) == 0 {
		return errors.New("recipe must have at least 1 ingredient")
	}

	r.data[recipe.Name] = recipe
	return nil
}

func (r *RecipeRepository) GetAll() (map[string]Recipe, error) {
	if len(r.data) == 0 {
		return nil, errors.New("recipe not found!")
	}
	return r.data, nil
}

func (r *RecipeRepository) SearchByCategory(category string) ([]Recipe, error) {

	if category == "" {
		return nil, errors.New("category name cannot be empty")
	}

	var result []Recipe

	for _, recipe := range r.data {
		if recipe.Category == category {
			result = append(result, recipe)
		}
	}

	return result, nil
}

func (r *RecipeRepository) DeleteRecipe(recipeName string) error {
	recipeName = strings.TrimSpace(recipeName)

	if recipeName == "" {
		return errors.New("recipe name is required")
	}

	for key, recipe := range r.data {
		if strings.EqualFold(recipe.Name, recipeName) {
			delete(r.data, key)
			return nil
		}
	}

	return errors.New("recipe not found")
}

func deleteRecipe(r *RecipeRepository) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter recipe name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if err := r.DeleteRecipe(name); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Recipe deleted successfully!")
}

// Business Logic
func MainMenu() {
	store := NewRecipeRepository()

	var choice int

	for {
		fmt.Println("1. Add Recipe")
		fmt.Println("2. View All Recipe")
		fmt.Println("3. Search by Category")
		fmt.Println("4. View Details")
		fmt.Println("5. Delete Recipe")
		fmt.Println("6. Exit")

		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			AddRecipe(store)
		case 2:
			ViewAllRecipes(store)
		case 3:
			SearchByCategory(store)
		case 4:
			getRecipeByRecipeName(store)
		case 5:
			deleteRecipe(store)
		case 6:
			fmt.Println("Bye")
			return

		default:
			fmt.Println("Invalid Choice")
		}
	}
}

func AddRecipe(store *RecipeRepository) {
	var name, category string
	var ingredient []string
	var steps []string
	var cookTime int

	fmt.Print("Enter Recipe Name: ")
	fmt.Scan(&name)

	fmt.Print("Enter Category Name: ")
	fmt.Scan(&category)

	// Input cooktime
	fmt.Println("Enter CookTime: ")
	fmt.Scan(&cookTime)

	// Input Ingredient
	for {
		var ing string

		fmt.Print("Enter ingredient (type 'done' to finish): ")
		fmt.Scan(&ing)

		if ing == "done" {
			break
		}

		ingredient = append(ingredient, ing)
	}

	// Input recipe step
	for {
		var step string

		fmt.Print("Enter Step (type 'done' to finish): ")
		fmt.Scan(&step)

		if step == "done" {
			break
		}

		steps = append(steps, "- "+step)
	}

	err := store.Add(Recipe{
		Name:        name,
		Category:    category,
		CookTime:    cookTime,
		Ingredients: ingredient,
		Steps:       steps,
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Recipe added successfully!")
}

func ViewAllRecipes(store *RecipeRepository) error {
	recipe, err := store.GetAll()

	if err != nil {
		return fmt.Errorf("add recipe: %w", err)
	}

	for key, value := range recipe {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}

	return nil

}

func SearchByCategory(store *RecipeRepository) {
	var categoryName string

	fmt.Print("Enter category: ")
	fmt.Scan(&categoryName)

	recipes, err := store.SearchByCategory(categoryName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(recipes) == 0 {
		fmt.Println("No recipes found.")
		return
	}

	for _, recipe := range recipes {
		fmt.Printf("Name: %s | Category: %s\n", recipe.Name, recipe.Category)
	}
}

func (r *RecipeRepository) GetByRecipeName(recipeName string) (Recipe, error) {
	recipeName = strings.TrimSpace(recipeName)

	if recipeName == "" {
		return Recipe{}, errors.New("recipe name is required")
	}

	for _, recipe := range r.data {
		if strings.EqualFold(recipe.Name, recipeName) {
			return recipe, nil
		}
	}

	return Recipe{}, errors.New("recipe not found")
}

func getRecipeByRecipeName(r *RecipeRepository) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter recipe name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	recipe, err := r.GetByRecipeName(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	printRecipe(recipe)
}

func printRecipe(recipe Recipe) {
	fmt.Println("┌─────────────────────────────────┐")
	fmt.Printf("│ Name     : %-20s │\n", recipe.Name)
	fmt.Printf("│ Category : %-20s │\n", recipe.Category)
	fmt.Printf("│ CookTime : %-17d mins │\n", recipe.CookTime)

	fmt.Println("├─────────────────────────────────┤")
	fmt.Println("│ Ingredients:                    │")
	for _, ingredient := range recipe.Ingredients {
		fmt.Printf("│  - %-28s │\n", ingredient)
	}

	fmt.Println("├─────────────────────────────────┤")
	fmt.Println("│ Steps:                          │")
	for i, step := range recipe.Steps {
		fmt.Printf("│  %d. %-27s │\n", i+1, step)
	}

	fmt.Println("└─────────────────────────────────┘")
}

// Entry Point
func main() {
	MainMenu()
}
