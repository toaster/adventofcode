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
	time, _ := strconv.Atoi(lines[0])
	var ids []int
	for _, i := range strings.Split(lines[1], ",") {
		id, _ := strconv.Atoi(i)
		if id == 0 {
			continue
		}
		ids = append(ids, id)
	}
	bestID := 0
	timeToWait := 0
	for _, id := range ids {
		t := id - time%id
		if timeToWait == 0 || timeToWait > t {
			timeToWait = t
			bestID = id
		}
	}
	fmt.Println("line", bestID, "* minutes to wait", timeToWait, "=", bestID*timeToWait)
}
