package utils

import (
	"reflect"
	"testing"
)

func TestFilterSlice(t *testing.T) {
	var params = []struct {
		description string
		inSlice     []int
		outSlice    []int
	}{
		{"nil slice", nil, []int{}},
		{"empty slice", []int{}, []int{}},
		{"slice with exclude elements", []int{1, 4, 5, 12, -1}, []int{5, -1}},
		{"slice without exclude elements", []int{-1, 3, 5, 17}, []int{-1, 3, 5, 17}},
	}
	for _, param := range params {
		result := FilterSlice(param.inSlice)
		if len(result) != len(param.outSlice) {
			t.Errorf("%s: invalid len, src slice %v, expected slice %v", param.description, result, param.outSlice)
		}
		if cap(result) != cap(param.outSlice) {
			t.Errorf("%s: invalid cap, src slice %v, expected slice %v", param.description, result, param.outSlice)
		}
		if !reflect.DeepEqual(result, param.outSlice) {
			t.Errorf("%s: expected slice '%v' not equal result slice '%v'", param.description, param.outSlice, result)
		}
	}
}
