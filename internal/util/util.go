package util

// CopyMap copies a map.
func CopyMap[K comparable, V any](value map[K]V) map[K]V {
	newValue := map[K]V{}
	for k, v := range value {
		newValue[k] = v
	}
	return newValue
}

// CopyMapExcept copies a map except a certain key.
func CopyMapExcept[K comparable, V any](value map[K]V, key K) map[K]V {
	newValue := map[K]V{}
	for k, v := range value {
		if k == key {
			continue
		}
		newValue[k] = v
	}
	return newValue
}

// SlicesEqual returns whether two slices contain the same elements in the same order.
func SlicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	return true
}
