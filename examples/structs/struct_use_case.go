package main

import "fmt"

// Define a few structs for the use cases

// Employee struct for representing an employee
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
}

// Rectangle struct for geometric calculations
type Rectangle struct {
	Width, Height float64
}

// Car struct for representing a car with attributes
type Car struct {
	Make, Model string
	Year        int
}

// Book struct for a library system
type Book struct {
	Title, Author string
	Pages         int
}

// User struct for a user profile in an application
type User struct {
	Username, Email string
	Age             int
}

func main() {
	// Use Case 1: Storing and Accessing Struct Data
	emp := Employee{1, "John", "Doe", "Engineer"}
	fmt.Println("Employee:", emp)

	// Use Case 2: Using Structs with Functions
	rect := Rectangle{Width: 10, Height: 5}
	area := rect.Width * rect.Height
	fmt.Println("Rectangle Area:", area)

	// Use Case 3: Array of Structs
	cars := [2]Car{
		{"Toyota", "Corolla", 2020},
		{"Honda", "Civic", 2021},
	}
	fmt.Println("Cars:", cars)

	// Use Case 4: Slice of Structs
	books := []Book{
		{"The Go Programming Language", "Alan A. A. Donovan", 380},
		{"Learning Go", "Miek Gieben", 300},
	}
	for _, book := range books {
		fmt.Printf("Book: %s by %s\n", book.Title, book.Author)
	}

	// Use Case 5: Maps with Struct Values
	users := make(map[int]User)
	users[101] = User{"johndoe", "johndoe@example.com", 30}
	fmt.Println("Users:", users)
}
