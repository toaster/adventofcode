package io

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ParseInts parses and returns a sequence of integer separated by given separator.
func ParseInts(line string, sep string) (numbers []int) {
	for _, s := range strings.Split(line, sep) {
		if s == "" {
			continue
		}

		v, err := strconv.Atoi(s)
		if err != nil {
			ReportError("failed to parse input", err)
		}

		numbers = append(numbers, v)
	}
	return
}

// ReadLines reads all lines from os.Stdin.
func ReadLines() []string {
	input, err := ioutil.ReadAll(os.Stdin)
	ReportError("failed to read standard input", err)
	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}

// ReportError reports an error and exits the program.
func ReportError(msg string, err error) {
	if err != nil {
		if msg != "" {
			_, _ = fmt.Fprintln(os.Stderr, msg+":", err)
		} else {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}
