package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	var ratings []int
	for _, line := range lines {
		rating, _ := strconv.Atoi(line)
		ratings = append(ratings, rating)
	}
	sort.Ints(ratings)
	ratings = append(ratings, ratings[len(ratings)-1]+3)

	segmentLength := 0
	mutations := 1
	for i, rating := range ratings {
		var diff int
		if i == 0 {
			diff = rating
		} else {
			diff = rating - ratings[i-1]
		}
		if diff == 1 {
			segmentLength++
			continue
		}
		if segmentLength > 0 {
			mutations *= possibleMutationCount(segmentLength)
			segmentLength = 0
		}
	}
	fmt.Println("mutations:", mutations)
}

func possibleMutationCount(segmentLength int) int {
	if segmentLength == 0 {
		return 1
	}
	if segmentLength == 1 {
		return 1
	}
	if segmentLength == 2 {
		return 2
	}
	if segmentLength == 3 {
		return 4
	}
	return 2*possibleMutationCount(segmentLength-1) - possibleMutationCount(segmentLength-4)
}
