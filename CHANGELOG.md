# Change Log

## [0.0.2] - 2022-04-27
### Added
- `Clone[T any](arr []T) []T`
- `Difference[T comparable](arr1 []T, arr2 []T) []T`
- `Exclusive[T comparable](arr1 []T, arr2 []T) []T`
- `Intersect[T comparable](arr1 []T, arr2 []T) []T`
- `RemoveAt[T any](arr []T, index int) ([]T, T)`
- `Replace[T any](arr []T, value T, pos ...int) []T`
- `Shuffle[T any](arr []T) []T`
- `Sum[T int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | uintptr | float32 | float64 | complex64 | complex128](arr []T) T`
- `Unique[T comparable](arr []T) []T`
### Changed
- Change generic parameter from `comparable` to `any` for function `At`, `ChangeEach`, `Chunk`, `Each`, `Every`, `Fill`, `Filter`, `FindWithCallback`, `FindLastWithCallback`, `FindIndexWithCallback`, `FindLastIndexWithCallback`, `Join`, `Pop`, `Reverse`, `Shift`, `Unshift`
