package gorray

import (
	"fmt"
	"log"
	"strings"
)

func slicingPos[T any](arr []T, start int, end int) (int, int) {
	arrLength := len(arr)
	if start < 0 {
		start = arrLength + start
	}
	if end < 0 {
		end = arrLength + end
	}
	return start, end
}

// Get the element from array at given `index`
// Allow negative index, it'll count back
func At[T any](arr []T, index int) T {
	if index < 0 {
		index = len(arr) + index
	}
	return arr[index]
}

// `length` must be greater than 1
func Chunk[T any](arr []T, length int) [][]T {
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
func Fill[T any](arr []T, value T, pos ...int) []T {
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

// Check if array has given elements
func Exist[T comparable](arr []T, elements ...T) bool {
	if len(arr) < 1 {
		return false
	}
	if len(elements) < 1 {
		return true
	}
	for _, element := range elements {
		notFound := true
		for _, elem := range arr {
			if elem == element {
				notFound = false
				break
			}
		}
		if notFound {
			return false
		}
	}
	return true
}

// Joins all element from array into one string.
//
// Not recommended because it'll joins all element with `fmt.Sprint`, use JoinStringer instead
func Join[T any](arr []T) string {
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
func Pop[T any](arr []T) ([]T, T) {
	pos := len(arr) - 1
	newArr := Clone(arr[0:pos])
	return newArr, arr[pos]
}

// Remove the first element on array, and return it.
func Shift[T any](arr []T) ([]T, T) {
	newArr := Clone(arr[1:])
	return newArr, arr[0]
}

// Add new element on array to the first position.
func Unshift[T any](arr []T, element T) []T {
	return append([]T{element}, arr...)
}

// Reverse the array
func Reverse[T any](arr []T) []T {
	newArr := Clone(arr)
	arrLength := len(arr)
	length := arrLength / 2
	if arrLength%2 > 0 {
		length++
	}
	arrLength -= 1
	for i := 0; i < length; i++ {
		if i == arrLength-i {
			continue
		}
		newArr[i] = arr[arrLength-i]
		newArr[arrLength-i] = arr[i]
	}
	return newArr
}

// Clone array tot the new array
func Clone[T any](arr []T) []T {
	newArr := make([]T, len(arr))
	copy(newArr, arr)
	return newArr
}

// Remove element from array at specified index.
//
// The element that removed will returned.
func RemoveAt[T any](arr []T, index int) ([]T, T) {
	newArray := Clone(arr[0:index])
	if index < len(arr) {
		newArray = append(newArray, arr[index+1:]...)
	}
	return newArray, arr[index]
}

// Replace element from array with given value at specified position
func Replace[T any](arr []T, value T, pos ...int) []T {
	arrLength := len(arr)
	var start, end int

	posLength := len(pos)
	if posLength > 0 {
		start = pos[0]
		end = start + 1
	}
	if posLength > 1 {
		end = pos[1]
	}
	start, end = slicingPos(arr, start, end)

	newArr := make([]T, arrLength)
	for index, elem := range newArr {
		if index >= start && index < end {
			newArr[index] = value
			continue
		}
		newArr[index] = elem
	}

	return newArr
}
