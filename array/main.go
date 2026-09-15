package main

import "fmt"

var globalArr = [3]string{"Hello", "World", "!"}

func main() {
	arr := [2]int{3, 6}

	fmt.Println(arr)

	fmt.Println(globalArr[1])
	fmt.Println(globalArr)
}

func init() {
	fmt.Println("This will be invoked first")
}
