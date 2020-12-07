package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nums []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	for _, num := range nums {
		if factor := findProduct(nums, 2020-num); factor != 0 {
			fmt.Println("Result:", num*factor)
			break
		}
	}
}

func findProduct(nums []int, sum int) int {
	neededs := map[int]*int{}
	for _, num := range nums {
		needed := sum - num
		neededs[num] = &needed
		if neededs[needed] != nil {
			return num * needed
		}
	}
	return 0
}
