package math

// AbsInt compute the absolute value of an integer.
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// MaxInt returns the maximum of two integers.
func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Sort2Int sorts exactly two integers.
func Sort2Int(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
