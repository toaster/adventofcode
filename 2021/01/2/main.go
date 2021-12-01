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
	reportErr("failed to read standard input", err)

	count := 0
	last := []int{0, 0, 0}
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		value, err := strconv.Atoi(line)
		reportErr("failed to parse input", err)

		if last[0] != 0 {
			if value+last[2]+last[1] > last[0]+last[1]+last[2] {
				count++
			}
		}
		last[0] = last[1]
		last[1] = last[2]
		last[2] = value
	}

	fmt.Println(count)
}

func reportErr(msg string, err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, msg+":", err)
		os.Exit(1)
	}
}
