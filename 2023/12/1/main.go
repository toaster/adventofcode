package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

const expansionFactor = 2

func main() {
	sum := 0
	for _, line := range io.ReadLines() {
		conditionRecordsAndDamagedCounts := strings.Split(line, " ")
		damagedCounts := io.ParseInts(conditionRecordsAndDamagedCounts[1], ",")
		sum += countPossibleArrangements(conditionRecordsAndDamagedCounts[0], damagedCounts)
	}
	fmt.Println(sum)
}

func countPossibleArrangements(conditionRecords string, damagedCounts []int) int {
	fmt.Println("records:", conditionRecords)
	rawPattern := strings.Builder{}
	rawPattern.WriteRune('^')
	for i, count := range damagedCounts {
		if i > 0 {
			rawPattern.WriteString("\\.+")
		}
		rawPattern.WriteString(fmt.Sprintf("#{%d}", count))
	}
	rawPattern.WriteRune('$')
	pattern := regexp.MustCompile(rawPattern.String())

	unsureCount := 0
	var unsureIndices []int
	for i, r := range conditionRecords {
		if r == '?' {
			unsureCount++
			unsureIndices = append(unsureIndices, i)
		}
	}
	candidateCount := 1 << unsureCount
	fmt.Println("unsure count:", unsureCount, "candidate count:", candidateCount)

	possibilitiesCount := 0
	for x := 0; x < candidateCount; x++ {
		candidate := []byte(conditionRecords)
		for i, index := range unsureIndices {
			if (x>>i)&1 == 1 {
				candidate[index] = '#'
			} else {
				candidate[index] = '.'
			}
		}
		if pattern.Match(bytes.Trim(candidate, ".")) {
			possibilitiesCount++
			fmt.Println("MATCH of candidate", x, ":", string(candidate))
		}
	}

	fmt.Println("==>", possibilitiesCount)
	return possibilitiesCount
}
