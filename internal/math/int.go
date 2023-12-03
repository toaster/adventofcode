package math

import "sort"

// AbsInt compute the absolute value of an integer.
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// AverageInt returns the (floor of the) average of a slice of integers.
func AverageInt(nums []int) int {
	count := len(nums)
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum / count
}

// ContainsInt returns whether a slice of integers contains a specific one.
func ContainsInt(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

// DetectDigit checks whether a byte shows a digit and returns its value if so.
func DetectDigit(c byte) (int, bool) {
	if c >= '0' && c <= '9' {
		return int(c - '0'), true
	}
	return 0, false
}

// MaxInt returns the maximum of two integers.
func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// MedianInt returns the (upper) median of a slice of integers.
func MedianInt(input []int) int {
	count := len(input)
	nums := make([]int, count)
	copy(nums, input)
	sort.Ints(nums)
	var m int
	if count%2 == 0 {
		m = nums[count/2]
	} else {
		m = nums[(count-1)/2]
	}
	return m
}

// MinInt returns the minimum of two integers.
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Sort2Int sorts exactly two integers.
func Sort2Int(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
