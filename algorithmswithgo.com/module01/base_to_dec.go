package module01

import (
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(digits string, base int) int {
	const BASE_DIGITS string = "0123456789ABCDEF"
	var result int
	for i, d := range digits {
		value := strings.Index(BASE_DIGITS, string(d))
		exponential := len(digits) - i - 1
		result += value * int(math.Pow(float64(base), float64(exponential)))
	}
	return result
}

func BaseToDecOptimiezed(digits string, base int) int {
	const BASE_DIGITS string = "0123456789ABCDEF"
	var result int
	power := 1
	for i := len(digits) - 1; i >= 0; i-- {
		value := strings.Index(BASE_DIGITS, string(digits[i]))
		result += value * power
		power *= base
	}
	return result
}
