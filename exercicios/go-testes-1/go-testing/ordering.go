package calc

import "sort"

func CrescentOrdering(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}
