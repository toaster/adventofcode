package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nums := map[int]*int{}
	var result *int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		needed := 2020 - num
		nums[num] = &needed
		if nums[needed] != nil {
			result = &num
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
	}
	fmt.Println("Result:", *result*(2020-*result))
}
