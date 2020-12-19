package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	sum := 0
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		v, _ := parseAndCompute(strings.ReplaceAll(line, " ", ""))
		sum += v
	}
	fmt.Println("sum:", sum)
}

func parseAndCompute(input string) (int, int) {
	result := 0
	op := uint8('+')
	for i := 0; i < len(input); i++ {
		token := input[i]
		switch token {
		case '(':
			value, skip := parseAndCompute(input[i+1:])
			i += skip
			i++ // closing brace
			result = compute(result, op, value)
		case ')':
			return result, i
		case '*':
			op = '*'
			value, skip := parseAndCompute(input[i+1:])
			i += skip
			result = compute(result, op, value)
		case '+':
			op = token
		default:
			result = compute(result, op, int(token-'0'))
		}
	}
	return result, len(input)
}

func compute(a int, op uint8, b int) int {
	switch op {
	case '+':
		return a + b
	case '*':
		return a * b
	}
	panic("oops")
}
