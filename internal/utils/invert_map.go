package utils

import "fmt"

func InvertStrStrMap(m map[string]string) (map[string]string, error) {
	result := make(map[string]string, len(m))
	for k, v := range m {
		if _, ok := result[v]; ok {
			return nil, fmt.Errorf("key '%s' already exists", v)
		}
		result[v] = k
	}
	return result, nil
}
