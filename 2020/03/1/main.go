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
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	patternWidth := len(lines[0])
	headingX := 3
	x := 0
	trees := 0
	for _, line := range lines {
		if line[x] == '#' {
			trees++
		}
		x += headingX
		x = x % patternWidth
	}
	fmt.Println("trees:", trees)
}
