package gorray

// Compare 2 arrays
// It also check if both array have a same length
func Compare[T comparable](arr1 []T, arr2 []T) bool {
	arr1Len := len(arr1)
	arr2Len := len(arr2)
	if arr1Len != arr2Len {
		return false
	}

	for index, elemArr1 := range arr1 {
		if elemArr1 != arr2[index] {
			return false
		}
	}

	return true
}
