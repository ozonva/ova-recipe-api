package utils

func SplitIntSlice(srcSlice []int, batchSize int) [][]int {
	if 0 == len(srcSlice) || 0 >= batchSize {
		return make([][]int, 0, 0)
	}
	subSlicesCount := len(srcSlice) / batchSize
	if 0 != len(srcSlice)%batchSize {
		subSlicesCount += 1
	}
	var result = make([][]int, 0, subSlicesCount)
	startIdx := 0
	for subSlicesCount > 1 {
		result = append(result, append(make([]int, 0, batchSize), srcSlice[startIdx:startIdx+batchSize]...))
		startIdx += batchSize
		subSlicesCount -= 1
	}
	result = append(result, append(make([]int, 0, len(srcSlice)-startIdx), srcSlice[startIdx:]...))
	return result
}
