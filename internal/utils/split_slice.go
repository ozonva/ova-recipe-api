package utils

import "ova-recipe-api/internal/recipe"

func min(lhs, rhs uint) uint {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func SplitIntSlice(srcSlice []int, batchSize uint) [][]int {
	if 0 == batchSize {
		return make([][]int, 0, 0)
	}
	srcSliceLen := uint(len(srcSlice))
	var result = make([][]int, 0, (srcSliceLen+batchSize-1)/batchSize)
	for startIdx := uint(0); startIdx < srcSliceLen; startIdx += batchSize {
		endIdx := min(startIdx+batchSize, srcSliceLen)
		subSlice := append(make([]int, 0, endIdx-startIdx), srcSlice[startIdx:endIdx]...)
		result = append(result, subSlice)
	}
	return result
}

func SplitRecipeSlice(srcSlice []recipe.Recipe, batchSize uint) [][]recipe.Recipe {
	if 0 == batchSize {
		return make([][]recipe.Recipe, 0, 0)
	}
	srcSliceLen := uint(len(srcSlice))
	var result = make([][]recipe.Recipe, 0, (srcSliceLen+batchSize-1)/batchSize)
	for startIdx := uint(0); startIdx < srcSliceLen; startIdx += batchSize {
		endIdx := min(startIdx+batchSize, srcSliceLen)
		subSlice := append(make([]recipe.Recipe, 0, endIdx-startIdx), srcSlice[startIdx:endIdx]...)
		result = append(result, subSlice)
	}
	return result
}