package main

import "fmt"

// Creates a copy of both varaibles recieved
func add(x int, y int) {
	x = y + x
}

// Saves sum on x address
func addPointers(x *int, y int) {
	*x = *x + y
}

func main() {
	x := 2
	y := 10

	// add(x, y)

	addPointers(&x, y)

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(y)
}
