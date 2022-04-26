package gorray

import (
	"fmt"
	"log"
	"strings"
)

// Get the element from array at given `index`
// Allow negative index, it'll count back
func At[T comparable](arr []T, index int) T {
	if index < 0 {
		index = len(arr) + index
	}
	return arr[index]
}

// `length` must be greater than 1
func Chunk[T comparable](arr []T, length int) [][]T {
	if length < 1 {
		panic("Length must be greater than 1")
	}

	arrLength := len(arr)
	chunksLength := arrLength / length
	if arrLength%length > 0 {
		chunksLength++
	}
	chunks := make([][]T, chunksLength)

	for i := 0; i < chunksLength; i++ {
		start := i * length
		end := start + length
		if end > arrLength {
			end = arrLength
		}
		chunks[i] = arr[start:end]
	}

	log.Println(len(chunks))
	return chunks
}

// Fill element to the array.
//
// You can specify which start and end position where element will be filled in Array.
func Fill[T comparable](arr []T, value T, pos ...int) []T {
	newArr := []T{}
	start := 0
	end := len(arr)
	if len(pos) > 0 {
		if len(pos) > 1 {
			if pos[1] > end {
				panic("end position must less than array's length")
			}
			end = pos[1]
		}
		start = pos[0]
		if start > end {
			panic("start position must less than array's length and end position")
		}
	}

	for index := start; index < end; index++ {
		newArr[index] = value
	}

	return arr
}

// Check if array has given element
func Exist[T comparable](arr []T, element T) bool {
	for _, elem := range arr {
		if elem == element {
			return true
		}
	}
	return false
}

// Joins all element from array into one string.
//
// Not recommended because it'll joins all element with `fmt.Sprint`, use JoinStringer instead
func Join[T comparable](arr []T) string {
	builder := new(strings.Builder)
	for _, element := range arr {
		builder.WriteString(fmt.Sprint(element))
	}
	return builder.String()
}

// Join all element from array Stringer into one string
func JoinStringer(arr []fmt.Stringer) string {
	builder := new(strings.Builder)
	for _, element := range arr {
		builder.WriteString(element.String())
	}
	return builder.String()
}

// Remove the last element on array, and return it.
func Pop[T comparable](arr []T) ([]T, T) {
	pos := len(arr) - 1
	newArr := arr[0:pos]
	return newArr, arr[pos]
}

// Remove the first element on array, and return it.
func Shift[T comparable](arr []T) ([]T, T) {
	newArr := arr[1:]
	return newArr, arr[0]
}

// Add new element on array to the first position.
func Unshift[T comparable](arr []T, element T) []T {
	return append([]T{element}, arr...)
}

// Reverse the array
func Reverse[T comparable](arr []T) []T {
	newArr := arr
	arrLength := len(arr)
	length := arrLength / 2
	if arrLength%2 > 0 {
		length++
	}
	for i := 0; i < length; i++ {
		if i == arrLength-i {
			continue
		}
		newArr[i] = arr[arrLength-i]
		newArr[arrLength-1] = arr[i]
	}
	return newArr
}
