package module01

import (
	"fmt"
	"strconv"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		n1, n2 := 3, 5
		var print string

		switch {
		case ModCheck(i, n1*n2):
			print = "Fizz Buzz"
		case ModCheck(i, n1):
			print = "Fizz"
		case ModCheck(i, n2):
			print = "Buzz"
		default:
			print = strconv.Itoa(i)
		}

		if i != n {
			print += ", "
			fmt.Print(print)
		}

		if i == n {
			fmt.Print(print)
			fmt.Println()
		}
	}
}

func ModCheck(n, mod int) bool {
	return n%mod == 0
}
