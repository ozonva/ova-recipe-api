package utils

import (
	"reflect"
	"testing"
)

func TestSplitIntSlice(t *testing.T) {
	var params = []struct {
		description string
		inSlice     []int
		batchSize   int
		outSlice    [][]int
	}{
		{"nil slice", nil, 0, [][]int{}},
		{"empty slice, batchSize = 0", []int{}, 0, [][]int{}},
		{"empty slice, batchSize < 0", []int{}, -1, [][]int{}},
		{"empty slice, batchSize > 0", []int{}, 1, [][]int{}},
		{"not empty slice, batchSize == 0", []int{1}, 0, [][]int{}},
		{"not empty slice, batchSize == len(slice)", []int{1}, 1, [][]int{{1}}},
		{"not empty slice, batchSize > len(slice)", []int{1}, 2, [][]int{{1}}},
		{"not empty slice, batchSize < len(slice)", []int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
		{"not empty slice (odd amount of elements), batchSize < len(slice)", []int{1, 2, 3, 4, 5}, 3, [][]int{{1, 2, 3}, {4, 5}}},
		{"not empty slice, batchSize == 1", []int{1, 2, 3, 4}, 1, [][]int{{1}, {2}, {3}, {4}}},
	}
	for _, param := range params {
		result := SplitIntSlice(param.inSlice, param.batchSize)
		if len(result) != len(param.outSlice) {
			t.Errorf("%s: invalid len, src slice %v, expected slice %v", param.description, result, param.outSlice)
		}
		if cap(result) != cap(param.outSlice) {
			t.Errorf("%s: invalid cap, src slice %v, expected slice %v", param.description, result, param.outSlice)
		}
		if !reflect.DeepEqual(result, param.outSlice) {
			t.Errorf("%s: resulting slice %v is not equal to expected %v", param.description, result, param.outSlice)
		}
	}
}
