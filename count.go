package gorray

// count all (total) values in arr
func CountAll[T comparable](arr []T) map[T]int {
	totals := make(map[T]int)

	for _, elem := range arr {
		totals[elem] += 1
	}

	return totals
}

// Calculate the sum of values in the array.
func Sum[
	T int8 | int16 | int32 | int64 | int |
		uint8 | uint16 | uint32 | uint64 | uint | uintptr |
		float32 | float64 |
		complex64 | complex128,
](arr []T) T {
	total := T(0)
	for _, elem := range arr {
		total += elem
	}
	return total
}
