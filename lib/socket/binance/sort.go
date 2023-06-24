package socket

import "sort"

func SortMap(streamTypeMap map[string]map[string]any, n int) []string {
	keys := make([]string, 0, n)

	for k = range streamTypeMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}
