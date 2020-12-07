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
		people := strings.Split(group, "\n")
		groupChoices := []rune(people[0])
		for _, choices := range people {
		GroupChoices:
			for i := 0; i < len(groupChoices); i++ {
				for _, choice := range choices {
					if choice == groupChoices[i] {
						continue GroupChoices
					}
				}
				if i+1 < len(groupChoices) {
					groupChoices = append(groupChoices[:i], groupChoices[i+1:]...)
				} else {
					groupChoices = groupChoices[:i]
				}
				i--
			}
		}
		sum += len(groupChoices)
	}
	fmt.Println("sum:", sum)
}
