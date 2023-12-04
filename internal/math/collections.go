package math

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// CommonElement2 returns one element that both slices have in common.
func CommonElement2[T constraints.Ordered](a, b []T) *T {
	sortedA := Sort(a)
	sortedB := Sort(b)
	j := 0
	for i := 0; i < len(sortedA); i++ {
		for sortedB[j] < sortedA[i] {
			j++
			if j == len(sortedB) {
				return nil
			}
		}
		if sortedB[j] == sortedA[i] {
			return &sortedB[j]
		}
	}
	return nil
}

// CommonElement3 returns one element that all three slices have in common.
func CommonElement3[T constraints.Ordered](a, b, c []T) *T {
	sortedA := Sort(a)
	sortedB := Sort(b)
	sortedC := Sort(c)
	j := 0
	k := 0
	for i := 0; i < len(sortedA); i++ {
		for sortedB[j] < sortedA[i] {
			j++
			if j == len(sortedB) {
				return nil
			}
		}
		for sortedC[k] < sortedA[i] {
			k++
			if k == len(sortedC) {
				return nil
			}
		}

		if sortedB[j] == sortedA[i] && sortedC[k] == sortedA[i] {
			return &sortedB[j]
		}
	}
	return nil
}

// Contains returns whether a slice contains a specific element.
func Contains[T comparable](nums []T, num T) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

// Intersection returns the intersection of two slices.
func Intersection[T comparable](a, b []T) (result []T) {
	for _, e := range a {
		if Contains(b, e) {
			result = append(result, e)
		}
	}
	return result
}

// Sort sorts a slice.
func Sort[T constraints.Ordered](input []T) []T {
	sorted := copySlice(input)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	return sorted
}

func copySlice[T any](input []T) []T {
	result := make([]T, len(input))
	copy(result, input)
	return result
}
