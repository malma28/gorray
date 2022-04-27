package gorray

// Check every element in array if passed true on callback
func Every[T any](arr []T, callback func(index int, element T) bool) bool {
	for index, element := range arr {
		if !callback(index, element) {
			return false
		}
	}
	return true
}

// Get every element in array that passed true on callback and insert that on new array
func Filter[T any](arr []T, callback func(index int, element T) bool) []T {
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
func FindWithCallback[T any](arr []T, callback func(index int, element T) bool) (T, bool) {
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
func FindIndexWithCallback[T any](arr []T, callback func(index int, element T) bool) int {
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
func FindLastWithCallback[T any](arr []T, callback func(index int, element T) bool) (T, bool) {
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
func FindLastIndexWithCallback[T any](arr []T, callback func(index int, element T) bool) int {
	arrLength := len(arr)
	for index := arrLength - 1; index >= 0; index++ {
		if callback(index, arr[index]) {
			return index
		}
	}
	return -1
}

// Pass all elements on array to callback
func Each[T any](arr []T, callback func(index int, element T)) {
	for index, element := range arr {
		callback(index, element)
	}
}

// Pass all elements on array to callback, but you can return the value
func ChangeEach[T any](arr []T, callback func(index int, element T) T) []T {
	newArr := make([]T, len(arr))
	for index, element := range arr {
		newArr[index] = callback(index, element)
	}
	return newArr
}

// Removes duplicate values from array.
func Unique[T comparable](arr []T) []T {
	newArr := []T{}
	for _, elem := range arr {
		if !Exist(newArr, elem) {
			newArr = append(newArr, elem)
		}
	}
	return newArr
}

// Select the element that only exist in both array
func Intersect[T comparable](arr1 []T, arr2 []T) []T {
	newArr := []T{}
	for _, elem1 := range arr1 {
		for _, elem2 := range arr2 {
			if elem1 == elem2 {
				newArr = append(newArr, elem1)
				break
			}
		}
	}
	return newArr
}

// Get elements array 1 that only exist array 1 but not in array 2
func Difference[T comparable](arr1 []T, arr2 []T) []T {
	newArr := []T{}
	arr1Cloned := Unique(arr1)

arr1Loop:
	for _, elem1 := range arr1Cloned {
		for _, elem2 := range arr2 {
			if elem2 == elem1 {
				continue arr1Loop
			}
		}
		newArr = append(newArr, elem1)
	}

	return newArr
}

// Get elements from both arrays that exist in only one of them and not in both.
func Exclusive[T comparable](arr1 []T, arr2 []T) []T {
	newArr := []T{}
	arr1Cloned := Unique(arr1)
	arr2Cloned := Unique(arr2)

arr1Loop:
	for _, elem1 := range arr1Cloned {
		for index, elem2 := range arr2Cloned {
			if elem2 == elem1 {
				arr2Cloned, _ = RemoveAt(arr2Cloned, index)
				continue arr1Loop
			}
		}
		newArr = append(newArr, elem1)
	}

	newArr = append(newArr, arr2Cloned...)

	return newArr
}
