package main

import (
	"fmt"
)

/*1
===========================================
  Go 100 Challenge
  Problem: #014
  Level: 🟡 Medium
===========================================

Problem: Hotel Room Booking System

Topic:
- struct
- custom type
- const
- slice
- functions
- status management

Industry Use:
Hotel Management Systems / Booking Platforms /
Airbnb / Property Management Systems

Rules (English):
- Manage hotel rooms with booking status
- Each room has: RoomNumber, Type, Price, Status
- Room types: Single, Double, Suite
- Status types: Available, Booked, Maintenance
- User can:
    * View all rooms
    * Book a room by room number
    * Checkout (set back to Available)
    * Set room to Maintenance
    * Filter rooms by status
    * Exit
- Already booked room   => "Room is already booked!"
- Room not found        => "Room not found!"
- Checkout empty room   => "Room is not booked!"

Rules (বাংলায়):
- Hotel rooms booking status সহ manage করবে
- প্রতিটা room এ থাকবে: RoomNumber, Type, Price, Status
- Room types: Single, Double, Suite
- Status: Available, Booked, Maintenance
- User করতে পারবে:
    * সব rooms দেখা
    * Room number দিয়ে book করা
    * Checkout করা (Available এ ফিরে যাবে)
    * Room কে Maintenance এ দেওয়া
    * Status দিয়ে filter করা
    * Exit করা
- Already booked room book করলে => "Room is already booked!"
- Room না পেলে                   => "Room not found!"
- Checkout খালি room করলে       => "Room is not booked!"

Example Run:
  === Hotel Booking System ===
  1. View All Rooms
  2. Book Room
  3. Checkout Room
  4. Set Maintenance
  5. Filter By Status
  6. Exit

  Choose: 1
  101  Single   500.00   Available
  102  Double   800.00   Available
  103  Suite   1500.00   Available

  Choose: 2
  Enter room number: 101
  Room 101 booked successfully!

  Choose: 1
  101  Single   500.00   Booked
  102  Double   800.00   Available
  103  Suite   1500.00   Available

  Choose: 5
  Filter by:
  1. Available
  2. Booked
  3. Maintenance
  Choose: 1
  102  Double   800.00   Available
  103  Suite   1500.00   Available
===========================================
*/

// types
type RoomStatus string

const (
	Available   RoomStatus = "Available"
	Booked      RoomStatus = "Booked"
	Maintenance RoomStatus = "Maintenance"
)

type RoomType string

const (
	Single RoomType = "Single"
	Double RoomType = "Double"
	Suite  RoomType = "Suite"
)

type Room struct {
	number   int
	roomType RoomType
	price    float64
	status   RoomStatus
}

func main() {

	// data
	rooms := &[]Room{
		{number: 101, roomType: Single, price: 500.00, status: Available},
		{number: 102, roomType: Double, price: 800.00, status: Available},
		{number: 103, roomType: Suite, price: 1500.00, status: Available},
		{number: 104, roomType: Single, price: 500.00, status: Available},
		{number: 105, roomType: Double, price: 800.00, status: Available},
	}
	MainMenu(rooms)
}

// main menu CLI
func MainMenu(rooms *[]Room) {

	var choice int

	for {
		fmt.Println()
		fmt.Println("======================================================")
		fmt.Println("                HOTEL ROOM MANAGEMENT")
		fmt.Println("======================================================")
		fmt.Println()
		fmt.Println("  [1] View All Rooms")
		fmt.Println("  [2] Book Room")
		fmt.Println("  [3] Checkout Room")
		fmt.Println("  [4] Set Room To Maintenance")
		fmt.Println("  [5] Filter Rooms By Status")
		fmt.Println()
		fmt.Println("  [0] Exit")
		fmt.Println()
		fmt.Println("------------------------------------------------------")
		fmt.Print(" Select Option > ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewAllRoom(rooms)
		case 2:
			bookRoom(rooms)
		case 3:
			checkoutRoom(rooms)
		case 0:
			return
		default:
			fmt.Println("Invlaid Input!")
		}

	}
}

// view all room
func viewAllRoom(rooms *[]Room) {
	for _, room := range *rooms {
		fmt.Printf(" %-8d %-10s %-10.2f %-10s\n",
			room.number,
			room.roomType,
			room.price,
			room.status,
		)
	}
}

func roomNumberInput() (int, error) {
	fmt.Println("Enter room number: ")
	var roomNumber int
	fmt.Scan(&roomNumber)

	if roomNumber <= 0 {
		return 0, fmt.Errorf("invalid room number")
	}

	if roomNumber == 3 {
		return 0, fmt.Errorf("room 3 is not allowed")
	}

	return roomNumber, nil
}

// Book room logic
func bookRoom(rooms *[]Room) {

	roomNumber, err := roomNumberInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range *rooms {

		if (*rooms)[i].number == roomNumber {
			if (*rooms)[i].status == Booked {
				fmt.Println("Room is already Booked!")
				return
			}

			if (*rooms)[i].status == Maintenance {
				fmt.Println("Room is under maintenance!")
				return
			}

			(*rooms)[i].status = Booked

			fmt.Printf("Room %d booked successfully!\n", roomNumber)
			return

		}
	}

	fmt.Println("Room not found!")

}

// checkout room logic
func checkoutRoom(rooms *[]Room) {
	roomNumber, err := roomNumberInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range *rooms {
		if (*rooms)[i].number == roomNumber {
			if !((*rooms)[i].status == Booked) {
				fmt.Println("❌ Room is not booked!")
				return
			}

			// checkout room book
			(*rooms)[i].status = Available

			fmt.Printf("✅ Room %d checkout successfully!\n", roomNumber)
			return
		}
	}

	fmt.Println("Room not found!")
}

// maintenance room
func maintenanceRoomBook(rooms *[]Room) {
	roomNumber, err := roomNumberInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range *rooms {
		if (*rooms)[i].number == roomNumber {

			if (*rooms)[i].status == Maintenance {
				fmt.Println("❌ Room already in Maintenance!")
				return
			}

			// checkout room book
			(*rooms)[i].status = Maintenance

			fmt.Printf("✅ Room %d checkout successfully!\n", roomNumber)
			return
		}
	}

	fmt.Println("Room not found!")
}

// filter by staus
func filterByStatus(rooms []Room) {

	for {
		fmt.Println("\nSelect Status:")
		fmt.Println("1. Available")
		fmt.Println("2. Booked")
		fmt.Println("3. Maintenance")
		fmt.Println("0. Back")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			filterByAvailableStatus(rooms)

		case 2:
			filterByBookedStatus(rooms)

		case 3:
			filterByMaintenanceStatus(rooms)

		case 0:
			fmt.Println("Goodbye! 👋")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func filterByAvailableStatus(rooms []Room) {
	fmt.Println("\n--- Available Rooms ---")

	for _, room := range rooms {
		if room.status == "Available" {
			fmt.Printf("%d %s %.2f %s\n",
				room.number,
				room.roomType,
				room.price,
				room.status,
			)
		}
	}
}

func filterByBookedStatus(rooms []Room) {
	fmt.Println("\n--- Booked Rooms ---")

	for _, room := range rooms {
		if room.status == "Booked" {
			fmt.Printf("%d %s %.2f %s\n",
				room.number,
				room.roomType,
				room.price,
				room.status,
			)
		}
	}
}

func filterByMaintenanceStatus(rooms []Room) {
	fmt.Println("\n--- Maintenance Rooms ---")

	for _, room := range rooms {
		if room.status == "Maintenance" {
			fmt.Printf("%d %s %.2f %s\n",
				room.number,
				room.roomType,
				room.price,
				room.status,
			)
		}
	}
}
