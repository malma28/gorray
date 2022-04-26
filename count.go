package gorray

// count all (total) values in arr
func CountAll[T comparable](arr []T) map[T]int {
	totals := make(map[T]int)

	for _, elem := range arr {
		totals[elem] += 1
	}

	return totals
}
