package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
	}

	rules := map[string]map[string]int{}
	for _, rule := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		colourAndContainments := strings.Split(rule, " bags contain ")
		colour := colourAndContainments[0]
		rules[colour] = parseContainments(colourAndContainments[1])
	}
	fmt.Println("sum:", countContainedBags("shiny gold", rules))
}

func countContainedBags(colour string, rules map[string]map[string]int) int {
	count := 0
	for c, i := range rules[colour] {
		count += i
		count += i * countContainedBags(c, rules)
	}
	return count
}

func parseContainments(value string) map[string]int {
	value = strings.ReplaceAll(value, " bags,", " bag,")
	value = strings.ReplaceAll(value, " bag.", "")
	value = strings.ReplaceAll(value, " bags.", "")
	containments := map[string]int{}
	for _, containment := range strings.Split(value, " bag, ") {
		countAndColour := strings.SplitN(containment, " ", 2)
		containments[countAndColour[1]], _ = strconv.Atoi(countAndColour[0])
	}
	return containments
}
