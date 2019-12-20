package math

// Sum up an integer array
func Sum(items []int) int {
	if items == nil {
		return 0
	}

	sum := 0
	for _, item := range items {
		sum = sum + item
	}

	return sum
}
