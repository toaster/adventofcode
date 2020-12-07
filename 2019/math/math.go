package math

func primeFactors(n int) map[int]int {
	pfs := map[int]int{}
	if n == 0 {
		return pfs
	}

	for n%2 == 0 {
		pfs[2]++
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs[i]++
			n = n / i
		}
	}

	if n > 2 {
		pfs[n]++
	}

	return pfs
}

// Abs compute the absolute value of an integer.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// LCM computes the lowest common multiplier of a slice of integers.
func LCM(numbers []int) int {
	lcmpfs := map[int]int{}
	for _, n := range numbers {
		pfs := primeFactors(n)
		for pf, c := range pfs {
			if c > lcmpfs[pf] {
				lcmpfs[pf] = c
			}
		}
	}
	lcm := 1
	for pf, c := range lcmpfs {
		for i := 0; i < c; i++ {
			lcm *= pf
		}
	}
	return lcm
}
