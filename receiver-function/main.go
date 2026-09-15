package main

import "fmt"

// User defines a custom struct type representing user profiles
type User struct {
	Name string
	Age  int
}

// 1. Regular Function Approach
// This function takes a User struct instance explicitly as a standard argument.
func printUserDetails(usr User) {
	fmt.Println("Regular Function Output:")
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

// 2. Receiver Function Approach (No extra parameters)
// The (usr User) block defines the receiver, binding this method to the User struct.
func (usr User) printDetails() {
	fmt.Println("Receiver Function Output:")
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

// 3. Receiver Function Approach (With an additional parameter)
// This binds to the User struct while accepting an extra integer argument 'a'.
func (usr User) call(a int) {
	fmt.Println("Receiver Function with Parameters Output:")
	fmt.Println("User Name:", usr.Name)
	fmt.Println("Passed Integer Value:", a)
}

func main() {
	// Initialize two distinct instances of the User struct
	user1 := User{
		Name: "Habib",
		Age:  30,
	}

	user2 := User{
		Name: "Roki",
		Age:  16,
	}

	// Executing the regular function approach
	printUserDetails(user1)
	printUserDetails(user2)

	fmt.Println("---------------------------------------")

	// Executing the receiver function approach using dot notation
	// The runtime automatically treats user1 as the receiver ('usr' inside the method)
	user1.printDetails()
	user2.printDetails()

	fmt.Println("---------------------------------------")

	// Executing a receiver function that handles extra parameter inputs
	// We pass the number 10 as an argument to the method
	user1.call(10)
}
