package main

import "fmt"

func main() {
	ageA := 32
	agePointer := &ageA
	fmt.Println("ageA", ageA)
	fmt.Println("agePointer", *agePointer)
	fmt.Println("ageA isAdult? ", isAdult(agePointer))
	yearsAsAdult(agePointer)
	fmt.Println("ageA yearsAsAdult", *agePointer)
	fmt.Println("ageA", ageA)

	ageB := 10
	yearsAsAdult(&ageB)
	fmt.Println("ageB yearsAsAdult", ageB)
}

const adultAge = 18

func isAdult(age *int) bool {
	return *age >= adultAge
}

func yearsAsAdult(age *int) {
	if isAdult(age) {
		*age = *age - adultAge
	} else {
		*age = 0
	}

}
