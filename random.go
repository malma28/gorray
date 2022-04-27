package gorray

import "math/rand"

// Randomizes the order of the elements in array.
//
// Before using this function it's better if you set a random seed if you don't want to get the same result.
func Shuffle[T any](arr []T) []T {
	newArr := Clone(arr)
	rand.Shuffle(len(newArr), func(i, j int) {
		newArr[i], newArr[j] = arr[j], arr[i]
	})
	return newArr
}
