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
	diffCounts := []int{0, 0, 0, 0}
	for i, rating := range ratings {
		if i == 0 {
			diffCounts[rating]++
		} else {
			diffCounts[rating-ratings[i-1]]++
		}
	}
	fmt.Println("diff counts:", diffCounts, "c(1)*c(3)", diffCounts[1]*diffCounts[3])
}
