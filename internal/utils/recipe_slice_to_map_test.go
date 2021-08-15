package utils

import (
	"fmt"
	"ova-recipe-api/internal/recipe"
	"reflect"
	"testing"
)

func TestRecipeSliceToMap(t *testing.T) {
	var params = []struct {
		description string
		inSlice     []recipe.Recipe
		outMap      map[uint64]recipe.Recipe
		err         error
	}{
		{"nil slice", nil, map[uint64]recipe.Recipe{}, nil},
		{"empty slice", []recipe.Recipe{}, map[uint64]recipe.Recipe{}, nil},
		{"slice w/o duplicates",
			[]recipe.Recipe{recipe.New(1, 1, "", "", nil), recipe.New(2, 1, "", "", nil)},
			map[uint64]recipe.Recipe{1: recipe.New(1, 1, "", "", nil), 2: recipe.New(2, 1, "", "", nil)},
			nil},
		{"slice with duplicates",
			[]recipe.Recipe{recipe.New(1, 1, "", "", nil), recipe.New(1, 1, "", "", nil)},
			nil,
			fmt.Errorf("duplicate recipe id '%d'", 1)},
	}
	for _, param := range params {
		result, err := RecipeSliceToMap(param.inSlice)
		if param.err != nil && param.err.Error() != err.Error() {
			t.Errorf("%s: expected error '%v' not equal result error '%v'", param.description, err, param.err)
		}
		if !reflect.DeepEqual(result, param.outMap) {
			t.Errorf("%s: resulting slice %v is not equal to expected %v", param.description, result, param.outMap)
		}
	}
}
