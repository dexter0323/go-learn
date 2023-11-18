package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	var total int
	for _, n := range numbers {
		total += n
	}
	return total
}
