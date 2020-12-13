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

	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	positionsByIDs := map[int]int{}
	var ids []int
	maxID := 0
	for i, rawID := range strings.Split(lines[1], ",") {
		id, _ := strconv.Atoi(rawID)
		if id == 0 {
			continue
		}
		if id > maxID {
			maxID = id
		}
		positionsByIDs[id] = i
		ids = append(ids, id)
	}
	x := 0
	p := ids[0]
	for i, id := range ids {
		if i == 0 {
			continue
		}
		x = findTime(x, p, id, positionsByIDs[id])
		p = p * id
	}
	fmt.Println("time:", x)
}

func findTime(start, p1, p2, delay int) int {
	x := start
	delay = delay % p2
	for ; ; x += p1 {
		if (p2-x%p2)%p2 != delay {
			continue
		}
		break
	}
	return x
}
