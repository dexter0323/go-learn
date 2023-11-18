package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList (list []int, n int) bool {
	for _, v := range list {
		if v == n {
			return true
		}
	}
	return false
}