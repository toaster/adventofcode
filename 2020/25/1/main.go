package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	cardPK, _ := strconv.Atoi(lines[0])
	doorPK, _ := strconv.Atoi(lines[1])
	cardLC := 0
	doorLC := 0
	cardValue := 1
	doorValue := 1
	subject := 7
	for doorValue != doorPK || cardValue != cardPK {
		if cardValue != cardPK {
			cardValue = transform(cardValue, subject, 1)
			cardLC++
		}
		if doorValue != doorPK {
			doorValue = transform(doorValue, subject, 1)
			doorLC++
		}
	}
	fmt.Println("cardKey:", transform(1, doorPK, cardLC), "doorKey:", transform(1, cardPK, doorLC))
}

func transform(v, s, c int) int {
	for i := 0; i < c; i++ {
		v = (v * s) % 20201227
	}
	return v
}
