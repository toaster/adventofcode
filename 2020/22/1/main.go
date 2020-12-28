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
		os.Exit(1)
	}

	var deckA []int
	var deckB []int
	for i, block := range strings.Split(strings.Trim(string(input), "\n"), "\n\n") {
		for _, line := range strings.Split(block, "\n") {
			if strings.HasPrefix(line, "Player") {
				continue
			}

			num, _ := strconv.Atoi(line)
			if i == 0 {
				deckA = append(deckA, num)
			} else {
				deckB = append(deckB, num)
			}
		}
	}
	for len(deckA) > 0 && len(deckB) > 0 {
		a := deckA[0]
		b := deckB[0]
		deckA = deckA[1:]
		deckB = deckB[1:]
		if a > b {
			deckA = append(deckA, a, b)
		} else {
			deckB = append(deckB, b, a)
		}
	}
	winnerDeck := deckA
	if len(deckA) == 0 {
		winnerDeck = deckB
	}
	score := calculateScore(winnerDeck)
	fmt.Println("winnerDeck score:", score)
}

func calculateScore(deck []int) (score int) {
	for i, num := range deck {
		score += num * (len(deck) - i)
	}
	return
}
