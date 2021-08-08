package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInvertMap(t *testing.T) {
	var params = []struct {
		description string
		inMap       map[string]string
		err         error
		outMap      map[string]string
	}{
		{"nil map", nil, nil, map[string]string{}},
		{"empty map", map[string]string{}, nil, map[string]string{}},
		{"valid map", map[string]string{"a": "1", "b": "2"}, nil, map[string]string{"1": "a", "2": "b"}},
		{"map with dup value", map[string]string{"a": "1", "b": "2", "c": "2"}, fmt.Errorf("key '2' already exists"), nil},
	}
	for _, param := range params {
		result, err := InvertStrStrMap(param.inMap)
		if param.err != nil && param.err.Error() != err.Error() {
			t.Errorf("%s: expected error '%v' not equal result error '%v'", param.description, err, param.err)
		}
		if !reflect.DeepEqual(result, param.outMap) {
			t.Errorf("%s: expected map '%v' not equal result map '%v'", param.description, param.outMap, result)
		}
	}
}
