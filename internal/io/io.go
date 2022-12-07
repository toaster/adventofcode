package io

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ParseBool parses a string into a boolean value.
func ParseBool(s string) bool {
	switch strings.ToLower(s) {
	case "on", "true", "yes", "1":
		return true
	}

	return false
}

// ParseInts parses and returns a sequence of integer separated by given separator.
func ParseInts(line string, sep string) (numbers []int) {
	for _, s := range strings.Split(line, sep) {
		if s == "" {
			continue
		}

		numbers = append(numbers, ParseInt(s))
	}
	return
}

// ParseInt parses and returns an integer.
func ParseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		ReportError("failed to parse input", err)
	}
	return v
}

// ReadFile reads all lines from a file.
func ReadFile(path string) []string {
	f, err := os.Open(path)
	ReportError("failed to open file", err)
	defer func() { _ = f.Close() }()
	input, err := ioutil.ReadAll(f)
	ReportError("failed to read file", err)
	return strings.Split(strings.Trim(string(input), "\n"), "\n")
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
