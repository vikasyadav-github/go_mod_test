package calc

import "errors"

// Given a range of numbers it returns the
// sum of those numbers
func Add(nums ...int) (error, int) {
	if len(nums) < 2 {
		return errors.New("There should be at least 2 arguments"), 0
	}
	total := 0
	for _, num := range nums {
		total += num
	}
	return nil, total
}
