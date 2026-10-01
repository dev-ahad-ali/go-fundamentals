package main

import (
	"fmt"
)

// User represents a custom struct blueprint.
// The struct name is Capitalized (Exported / Public).
type User struct {
	// Exported (Public) fields: Accessible from any package.
	Name string
	Age  int

	// Unexported (Private) field: Accessible ONLY inside the 'main' package.
	// External packages cannot access or modify this field directly.
	password string
}

// NewUser serves as a constructor function to initialize a User
// with both public and private properties safely.
func NewUser(name string, age int, rawPassword string) User {
	return User{
		Name:     name,
		Age:      age,
		password: rawPassword, // Initialized internally
	}
}

// GetPassword provides controlled read-only access to the unexported property.
func (u User) GetPassword() string {
	return u.password
}

// SetPassword provides encapsulated modification of the private property.
// Uses a pointer receiver (*User) to modify the actual memory block.
func (u *User) SetPassword(newPassword string) {
	if len(newPassword) >= 8 {
		u.password = newPassword
	}
}

func main() {
	// 1. Instantiating a struct directly (Literal Instantiation)
	user1 := User{
		Name:     "Habib",
		Age:      30,
		password: "mySecretPassword123", // Allowed here since main is in the same package
	}

	// 2. Instantiating via Constructor Function
	user2 := NewUser("Ahad", 25, "initialPass456")

	// 3. Accessing Public and Private fields via Dot Notation
	fmt.Println("--- Initial State ---")
	fmt.Println("User 1 Name:", user1.Name)
	fmt.Println("User 1 Age:", user1.Age)
	fmt.Println("User 1 Private Password:", user1.password)

	// 4. Demonstrating Memory Separation (Value Copies)
	// Assigning user1 to user3 creates an independent copy in memory.
	user3 := user1
	user3.Name = "Habib Updated"

	fmt.Println("\n--- Memory Isolation Demonstration ---")
	fmt.Println("User 1 Name (Original):", user1.Name) // Remains "Habib"
	fmt.Println("User 3 Name (Copy):    ", user3.Name) // Modified to "Habib Updated"

	// 5. Updating Private Field via Encapsulated Method
	fmt.Println("\n--- Encapsulation & Pointer Method ---")
	fmt.Println("User 2 Password before update:", user2.GetPassword())

	user2.SetPassword("newStrongPassword789")
	fmt.Println("User 2 Password after update: ", user2.GetPassword())
}
