package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s <PREAMBLE_LENGTH>\n", os.Args[0])
		os.Exit(1)
	}
	preambleLength, err := strconv.Atoi(os.Args[1])
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "preamble length", os.Args[1], "must be an integer:", err)
		os.Exit(1)
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	var nums []int
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}
	preamble := nums[:preambleLength]
	nums = nums[preambleLength:]

	invalid := 0
	invalid = findInvalid(nums, preamble)
	fmt.Println("first invalid:", invalid)
}

func findInvalid(nums []int, preamble []int) int {
	for _, num := range nums {
		if !isValid(num, preamble) {
			return num
		}
		preamble = append(preamble[1:], num)
	}
	return 0
}

func isValid(num int, preamble []int) bool {
	for i, p := range preamble {
		if p < num {
			for _, q := range preamble[i+1:] {
				if p+q == num {
					return true
				}
			}
		}
	}
	return false
}
