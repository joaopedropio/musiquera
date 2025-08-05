package utils

import "slices"

func Unique[T comparable](input []T) []T {
	seen := make(map[T]struct{})
	result := input[:0]
	for _, v := range input {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

func UniqueBy[T any, K comparable](input []T, keySelector func(T) K) []T {
	seen := make(map[K]struct{})
	result := input[:0]

	for _, v := range input {
		key := keySelector(v)
		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

func Find[T any](s []T, predicate func(T) bool) (T, bool) {
	index := slices.IndexFunc(s, predicate)
	if index >= 0 {
		return s[index], true
	}
	var zero T
	return zero, false
}
