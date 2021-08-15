package utils

import (
	"ova-recipe-api/internal/recipe"
	"reflect"
	"testing"
)

func TestSplitIntSlice(t *testing.T) {
	var params = []struct {
		description string
		inSlice     []int
		batchSize   uint
		outSlice    [][]int
	}{
		{"nil slice", nil, 0, [][]int{}},
		{"empty slice, batchSize = 0", []int{}, 0, [][]int{}},
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

func TestSplitRecipeSlice(t *testing.T) {
	var params = []struct {
		description string
		inSlice     []recipe.Recipe
		batchSize   uint
		outSlice    [][]recipe.Recipe
	}{
		{"nil slice", nil, 0, [][]recipe.Recipe{}},
		{"empty slice, batchSize = 0", []recipe.Recipe{}, 0, [][]recipe.Recipe{}},
		{"empty slice, batchSize > 0", []recipe.Recipe{}, 1, [][]recipe.Recipe{}},
		{"not empty slice, batchSize == 0", []recipe.Recipe{recipe.New(1, 1, "", "", nil)}, 0, [][]recipe.Recipe{}},
		{"not empty slice, batchSize == len(slice)", []recipe.Recipe{recipe.New(1, 1, "", "", nil)}, 1, [][]recipe.Recipe{{recipe.New(1, 1, "", "", nil)}}},
		{"not empty slice, batchSize > len(slice)", []recipe.Recipe{recipe.New(1, 1, "", "", nil)}, 2, [][]recipe.Recipe{{recipe.New(1, 1, "", "", nil)}}},
		{"not empty slice, batchSize < len(slice)",
			[]recipe.Recipe{recipe.New(1, 1, "", "", nil), recipe.New(2, 2, "", "", nil), recipe.New(3, 3, "", "", nil), recipe.New(4, 4, "", "", nil)}, 2,
			[][]recipe.Recipe{{recipe.New(1, 1, "", "", nil), recipe.New(2, 2, "", "", nil)}, {recipe.New(3, 3, "", "", nil), recipe.New(4, 4, "", "", nil)}}},
		{"not empty slice (odd amount of elements), batchSize < len(slice)",
			[]recipe.Recipe{recipe.New(1, 1, "", "", nil), recipe.New(2, 2, "", "", nil), recipe.New(3, 3, "", "", nil), recipe.New(4, 4, "", "", nil), recipe.New(5, 5, "", "", nil)}, 3,
			[][]recipe.Recipe{{recipe.New(1, 1, "", "", nil), recipe.New(2, 2, "", "", nil), recipe.New(3, 3, "", "", nil)}, {recipe.New(4, 4, "", "", nil), recipe.New(5, 5, "", "", nil)}}},
		{"not empty slice, batchSize == 1",
			[]recipe.Recipe{recipe.New(1, 1, "", "", nil), recipe.New(2, 2, "", "", nil), recipe.New(3, 3, "", "", nil), recipe.New(4, 4, "", "", nil)}, 1,
			[][]recipe.Recipe{{recipe.New(1, 1, "", "", nil)}, {recipe.New(2, 2, "", "", nil)}, {recipe.New(3, 3, "", "", nil)}, {recipe.New(4, 4, "", "", nil)}}},
	}
	for _, param := range params {
		result := SplitRecipeSlice(param.inSlice, param.batchSize)
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