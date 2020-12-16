package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	action rune
	amount int
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	start := strings.Split(strings.Trim(string(input), "\n"), ",")
	positionsByNums := map[int][]int{}
	lastNum := 0
	for i, r := range start {
		lastNum, _ = strconv.Atoi(r)
		positionsByNums[lastNum] = append(positionsByNums[lastNum], i+1)
	}
	for i := len(positionsByNums); i < 2020; i++ {
		p := positionsByNums[lastNum]
		if len(p) < 2 {
			lastNum = 0
		} else {
			lastNum = p[len(p)-1] - p[len(p)-2]
		}
		positionsByNums[lastNum] = append(positionsByNums[lastNum], i+1)
	}
	fmt.Println("2020th number:", lastNum)
}
