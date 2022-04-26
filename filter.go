package gorray

// Check every element in array if passed true on callback
func Every[T comparable](arr []T, callback func(index int, element T) bool) bool {
	for index, element := range arr {
		if !callback(index, element) {
			return false
		}
	}
	return true
}

// Get every element in array that passed true on callback and insert that on new array
func Filter[T comparable](arr []T, callback func(index int, element T) bool) []T {
	newArr := []T{}
	for index, element := range arr {
		if callback(index, element) {
			newArr = append(newArr, element)
		}
	}
	return newArr
}

// Return the first element that passed true on callback.
//
// Return false on second returned value if not found.
func FindWithCallback[T comparable](arr []T, callback func(index int, element T) bool) (T, bool) {
	for index, element := range arr {
		if callback(index, element) {
			return element, true
		}
	}
	var notFound T
	return notFound, false
}

// Return the first element's index that passed true on callback.
//
// Return -1 if not found.
func FindIndexWithCallback[T comparable](arr []T, callback func(index int, element T) bool) int {
	for index, element := range arr {
		if callback(index, element) {
			return index
		}
	}
	return -1
}

// Return the last element that passed true on callback.
//
// Return false on second returned value if not found.
func FindLastWithCallback[T comparable](arr []T, callback func(index int, element T) bool) (T, bool) {
	arrLength := len(arr)
	for index := arrLength - 1; index >= 0; index++ {
		if callback(index, arr[index]) {
			return arr[index], true
		}
	}
	var notFound T
	return notFound, false
}

// Return the last element's index that passed true on callback.
//
// Return -1 if not found.
func FindLastIndexWithCallback[T comparable](arr []T, callback func(index int, element T) bool) int {
	arrLength := len(arr)
	for index := arrLength - 1; index >= 0; index++ {
		if callback(index, arr[index]) {
			return index
		}
	}
	return -1
}

// Pass all elements on array to callback
func Each[T comparable](arr []T, callback func(index int, element T)) {
	for index, element := range arr {
		callback(index, element)
	}
}

// Pass all elements on array to callback, but you can return the value
func ChangeEach[T comparable](arr []T, callback func(index int, element T) T) []T {
	newArr := make([]T, len(arr))
	for index, element := range arr {
		newArr[index] = callback(index, element)
	}
	return newArr
}
