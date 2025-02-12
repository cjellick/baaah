package typed

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Concat[K constraints.Ordered, V any](maps ...map[K]V) map[K]V {
	result := map[K]V{}
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func SortedKeys[K constraints.Ordered, T any](data map[K]T) (result []K) {
	for k := range data {
		result = append(result, k)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return
}

type Entry[K, V any] struct {
	Key   K
	Value V
}

func Sorted[K constraints.Ordered, V any](data map[K]V) []Entry[K, V] {
	var result []Entry[K, V]
	for _, key := range SortedKeys(data) {
		result = append(result, Entry[K, V]{
			Key:   key,
			Value: data[key],
		})
	}
	return result
}
