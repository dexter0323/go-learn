package main

import "testing"

func TestSu(t *testing.T) {
	total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Sum is incorrect, got %d expected %d", total, 10)
	}
}

// go test -coverprofile=coverage.out
// go tool cover -func=coverage.out
// go tool cover -html=coverage.out

func TestFib(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}
	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", fib, item.n)
		}
	}
}

// go test -cpuprofile=profile.out
// go tool pprof profile.out -> top -> list Fibonacci -> web
