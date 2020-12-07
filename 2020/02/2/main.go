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
		letter := ruleComponents[1][0]
		ruleRangeComponents := strings.Split(ruleRange, "-")
		a, _ := strconv.Atoi(ruleRangeComponents[0])
		b, _ := strconv.Atoi(ruleRangeComponents[1])

		if (pass[a-1] == letter) != (pass[b-1] == letter) {
			validCount++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
	}
	fmt.Println("valid passwords:", validCount)
}
