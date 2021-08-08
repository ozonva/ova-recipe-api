package utils

func SplitIntSlice(srcSlice []int, batchSize uint) [][]int {
	srcSliceLen := uint(len(srcSlice))
	if 0 == srcSliceLen || 0 == batchSize {
		return make([][]int, 0, 0)
	}
	subSlicesCount := srcSliceLen / batchSize
	if 0 != srcSliceLen%batchSize {
		subSlicesCount += 1
	}
	var result = make([][]int, 0, subSlicesCount)
	startIdx := uint(0)
	for subSlicesCount > 1 {
		result = append(result, append(make([]int, 0, batchSize), srcSlice[startIdx:startIdx+batchSize]...))
		startIdx += batchSize
		subSlicesCount -= 1
	}
	result = append(result, append(make([]int, 0, srcSliceLen-startIdx), srcSlice[startIdx:]...))
	return result
}
