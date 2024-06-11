package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go doSomething(c)
	<-c
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done!")
	c <- 1
}

func Sum(x, y int) int {
	return x + y
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
