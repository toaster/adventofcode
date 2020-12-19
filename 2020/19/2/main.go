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

	// rule 0 is always "8 11" and the only user of 8 and 11
	// 8: 42 | 42 8
	prefixExp := regexp.MustCompile("^" + expandedRule(rawRules, rules, "42") + "(?P<remainder>.*)$")
	// 11: 42 31 | 42 11 31
	remainderExp := regexp.MustCompile("^" + expandedRule(rawRules, rules, "42") + "(?P<substring>.*)" + expandedRule(rawRules, rules, "31") + "$")
	sum := 0
	for _, msg := range strings.Split(rulesAndMsgs[1], "\n") {
		if matches(prefixExp, remainderExp, msg) {
			sum++
		}
	}
	fmt.Println("sum:", sum)
}

func expandedRule(rawRules, rules map[string]string, key string) string {
	if rules[key] != "" {
		return rules[key]
	}
	if rawRules[key] == "" {
		return ""
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
		if t == "|" || t[0] == '+' || t[0] == '(' || t[0] == ')' {
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

func matches(prefixExp, remainderExp *regexp.Regexp, msg string) bool {
	if prefixExp != nil {
		m := prefixExp.FindStringSubmatch(msg)
		if m == nil {
			return false
		}
		remainder := m[prefixExp.SubexpIndex("remainder")]
		if remainder == "" {
			return false
		}
		if matches(prefixExp, remainderExp, remainder) {
			return true
		}
		return matches(nil, remainderExp, remainder)
	}

	m := remainderExp.FindStringSubmatch(msg)
	if m == nil {
		return false
	}

	substr := m[remainderExp.SubexpIndex("substring")]
	if substr == "" {
		return true
	}

	return matches(nil, remainderExp, substr)
}
