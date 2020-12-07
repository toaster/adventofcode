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
			if !isFieldValid(field, entries[field]) {
				continue Passports
			}
		}
		valid++
	}
	fmt.Println("valid:", valid)
}

func isFieldValid(field, value string) bool {
	if value == "" {
		return false
	}
	switch field {
	case "byr":
		return isValidNum(value, 1920, 2002)
	case "iyr":
		return isValidNum(value, 2010, 2020)
	case "eyr":
		return isValidNum(value, 2020, 2030)
	case "hgt":
		return isValidLength(value, 150, 193, 59, 76)
	case "hcl":
		return isValidColour(value)
	case "ecl":
		return isValidColourName(value)
	case "pid":
		return isValidID(value, 9)
	}
	return true
}

func isValidColour(value string) bool {
	if len(value) != 7 {
		return false
	}
	if value[0] != '#' {
		return false
	}
	for _, c := range value[1:] {
		if c >= '0' && c <= '9' {
			continue
		}
		if c >= 'a' && c <= 'f' {
			continue
		}
		return false
	}
	return true
}

func isValidColourName(value string) bool {
	switch value {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

func isValidID(value string, length int) bool {
	if len(value) != length {
		return false
	}
	for _, c := range value {
		if c >= '0' && c <= '9' {
			continue
		}
		return false
	}
	return true
}

func isValidLength(value string, cmMin, cmMax, inMin, inMax int) bool {
	unit := value[len(value)-2:]
	value = value[:len(value)-2]
	switch unit {
	case "cm":
		return isValidNum(value, cmMin, cmMax)
	case "in":
		return isValidNum(value, inMin, inMax)
	}
	return false
}

func isValidNum(value string, min, max int) bool {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	if strconv.Itoa(intValue) != value {
		return false
	}
	return intValue >= min && intValue <= max
}
