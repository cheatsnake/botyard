package extlib

func SliceFilter[T any](slice []T, cap int, predicate func(T) bool) []T {
	filteredSlice := make([]T, 0, cap)

	for _, element := range slice {
		if predicate(element) {
			filteredSlice = append(filteredSlice, element)
		}
	}

	return filteredSlice
}

func SliceReverse[T any](slice []T) []T {
	length := len(slice)
	reversed := make([]T, length)

	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		reversed[i] = slice[j]
	}

	return reversed
}

func SlicePaginate[T any](slice []T, page, limit int) []T {
	startIndex := (page - 1) * limit
	endIndex := page * limit

	if startIndex >= len(slice) {
		return nil
	}

	if endIndex > len(slice) {
		endIndex = len(slice)
	}

	returnSlice := make([]T, endIndex-startIndex)
	copy(returnSlice, slice[startIndex:endIndex])

	return returnSlice
}

func SliceRemoveElement[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}

	result := make([]T, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)

	return result
}
