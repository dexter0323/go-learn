package main

import "fmt"

func main() {
	number := 5
	double := createTransformer(2)(number)
	triple := createTransformer(3)(number)

	fmt.Println(double, triple)

	f := factorial(5)

	fmt.Println(f)

	v := variadic(1, 2, 3, 4, 5)

	fmt.Println(v)

}

func createTransformer(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func variadic(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}
