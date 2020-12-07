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
	groups := strings.Split(strings.Trim(string(input), "\n"), "\n\n")
	sum := 0
	for _, group := range groups {
		groupChoices := map[rune]bool{}
		for _, choices := range strings.Split(group, "\n") {
			for _, choice := range choices {
				groupChoices[choice] = true
			}
		}
		sum += len(groupChoices)
	}
	fmt.Println("sum:", sum)
}
