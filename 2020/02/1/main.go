package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	validCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		components := strings.Split(line, ": ")
		rule := components[0]
		pass := components[1]
		ruleComponents := strings.Split(rule, " ")
		ruleRange := ruleComponents[0]
		letter := rune(ruleComponents[1][0])
		ruleRangeComponents := strings.Split(ruleRange, "-")
		min, _ := strconv.Atoi(ruleRangeComponents[0])
		max, _ := strconv.Atoi(ruleRangeComponents[1])

		count := 0
		for _, l := range pass {
			if l == letter {
				count++
				if count > max {
					break
				}
			}
		}
		if count >= min && count <= max {
			validCount++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
	}
	fmt.Println("valid passwords:", validCount)
}
