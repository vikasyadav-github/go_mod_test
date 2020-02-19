package calc

// Given a range of numbers it returns the
// sum of those numbers
func Add(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
