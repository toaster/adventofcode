package io

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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
