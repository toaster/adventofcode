package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := []byte(os.Args[1])
	pattern := make([]int, len(input))
	for i, b := range input {
		pattern[i] = int(b - '0')
	}
	e1 := 0
	e2 := 1
	recipes := []int{3, 7}
	for {
		s := recipes[e1] + recipes[e2]
		if s > 9 {
			recipes = append(recipes, 1)
			if len(recipes) > 5 && equal(recipes[len(recipes)-6:], pattern) {
				break
			}
		}
		recipes = append(recipes, s%10)
		if len(recipes) > 5 && equal(recipes[len(recipes)-6:], pattern) {
			break
		}
		e1 = (e1 + recipes[e1] + 1) % len(recipes)
		e2 = (e2 + recipes[e2] + 1) % len(recipes)
	}
	fmt.Println(len(recipes) - len(pattern))
}
func mainPart1() {
	iterCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	e1 := 0
	e2 := 1
	recipes := []int{3, 7}
	for i := 0; i < iterCount+10; i++ {
		s := recipes[e1] + recipes[e2]
		if s > 9 {
			recipes = append(recipes, 1)
		}
		recipes = append(recipes, s%10)
		e1 = (e1 + recipes[e1] + 1) % len(recipes)
		e2 = (e2 + recipes[e2] + 1) % len(recipes)
	}
	fmt.Println(recipes[iterCount : iterCount+10])
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, va := range a {
		if va != b[i] {
			return false
		}
	}
	return true
}
