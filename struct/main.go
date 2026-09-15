package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var user1 User

	user1 = User{
		Name: "Habib",
		Age:  30,
	}

	user2 := User{ // instance or object
		Name: "User second",
		Age:  50,
	}

	fmt.Println("Name", user1.Name)
	fmt.Println("Age", user1.Age)

	fmt.Println("Name", user2.Name)
	fmt.Println("Age", user2.Age)
}

func init() {
	fmt.Println("This will be invoked first")
}
