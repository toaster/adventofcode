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
	}

	reverseRules := map[string][]string{}
	for _, rule := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		colourAndContainments := strings.Split(rule, " bags contain ")
		for _, containment := range parseContainments(colourAndContainments[1]) {
			reverseRules[containment] = append(reverseRules[containment], colourAndContainments[0])
		}
	}
	coloursContainingShinyGold := map[string]bool{"shiny gold": true}
	for done := false; !done; {
		done = true
		for colour := range coloursContainingShinyGold {
			for _, cc := range reverseRules[colour] {
				if !coloursContainingShinyGold[cc] {
					done = false
					coloursContainingShinyGold[cc] = true
				}
			}
		}
	}
	fmt.Println("sum:", len(coloursContainingShinyGold)-1)
}

func parseContainments(value string) []string {
	value = strings.ReplaceAll(value, " bags,", " bag,")
	value = strings.ReplaceAll(value, " bag.", "")
	value = strings.ReplaceAll(value, " bags.", "")
	var containments []string
	for _, containment := range strings.Split(value, " bag, ") {
		countAndColour := strings.SplitN(containment, " ", 2)
		containments = append(containments, countAndColour[1])
	}
	return containments
}
