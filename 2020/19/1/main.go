package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	rulesAndMsgs := strings.Split(strings.Trim(string(input), "\n"), "\n\n")
	rawRules := map[string]string{}
	rules := map[string]string{}
	for _, rule := range strings.Split(rulesAndMsgs[0], "\n") {
		numAndSpec := strings.Split(rule, ": ")
		rawRules[numAndSpec[0]] = numAndSpec[1]
	}

	exp := regexp.MustCompile("^" + expandedRule(rawRules, rules, "0") + "$")
	sum := 0
	for _, msg := range strings.Split(rulesAndMsgs[1], "\n") {
		if exp.Match([]byte(msg)) {
			sum++
		}
	}
	fmt.Println("rule:", expandedRule(rawRules, rules, "0"), "sum:", sum)
}

func expandedRule(rawRules, rules map[string]string, key string) string {
	if rules[key] != "" {
		return rules[key]
	}

	rawRule := rawRules[key]
	if rawRule[0] == '"' {
		rules[key] = rawRule[1:2]
		return rules[key]
	}

	var parsedTokens []string
	if strings.Contains(rawRule, "|") {
		parsedTokens = append(parsedTokens, "(")
	}
	for _, t := range strings.Split(rawRule, " ") {
		if t == "|" {
			parsedTokens = append(parsedTokens, t)
			continue
		}
		parsedTokens = append(parsedTokens, expandedRule(rawRules, rules, t))
	}
	if strings.Contains(rawRule, "|") {
		parsedTokens = append(parsedTokens, ")")
	}
	rules[key] = strings.Join(parsedTokens, "")
	return rules[key]
}
