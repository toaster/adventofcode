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
	passports := strings.Split(strings.Trim(string(input), "\n"), "\n\n")
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	valid := 0
Passports:
	for _, passport := range passports {
		rawEntries := strings.Split(strings.ReplaceAll(passport, "\n", " "), " ")
		entries := map[string]string{}
		for _, entry := range rawEntries {
			data := strings.Split(entry, ":")
			entries[data[0]] = data[1]
		}
		for _, field := range requiredFields {
			if entries[field] == "" {
				continue Passports
			}
		}
		valid++
	}
	fmt.Println("valid:", valid)
}
