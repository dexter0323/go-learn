package main

import "fmt"

func main() {
	x := 10
	p := &x         // p is a pointer to x
	fmt.Println(*p) // Dereference the pointer to get the value
	*p = 20         // Change the value through the pointer
	fmt.Println(x)  // x is now 20
}
