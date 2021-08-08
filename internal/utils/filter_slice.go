package utils

var excludeElements = makeSetExcludeElements([]int{1, 4, 8, 12, 6, 0})

func makeSetExcludeElements(s []int) map[int]bool {
	result := make(map[int]bool, len(s))
	for _, elem := range s {
		result[elem] = true
	}
	return result
}

func iterateOverIncludeSliceElements (s []int, f func(elem int))  {
	for _, elem := range s {
		if _, ok := excludeElements[elem]; !ok {
			f(elem)
		}
	}
}

func FilterSlice(s []int) []int {
	elemsCount := 0
	// Trying to avoid unnecessary memory allocations
	iterateOverIncludeSliceElements(s, func(elem int) { elemsCount += 1 })
	result := make([]int, 0, elemsCount)
	iterateOverIncludeSliceElements(s, func(elem int) { result = append(result, elem) })
	return result
}
