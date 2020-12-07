package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type playedMarble struct {
	value      int
	prev, next *playedMarble
}

var inputRegex *regexp.Regexp

func init() {
	inputRegex = regexp.MustCompile("(\\d+) players; last marble is worth (\\d+) points")
}

func main() {
	input := os.Args[1]
	matches := inputRegex.FindStringSubmatch(input)
	players, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	lastMarble, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	fmt.Println("winner's score: ", winnerScore(players, lastMarble))
}

func winnerScore(players, lastValue int) int {
	cur := &playedMarble{0, nil, nil}
	cur.next = cur
	cur.prev = cur
	scores := make([]int, players)
	player := -1
	for nextValue := 1; nextValue < lastValue; nextValue++ {
		player = (player + 1) % players
		if nextValue%23 == 0 {
			scores[player] += nextValue
			marble := cur.prev.prev.prev.prev.prev.prev.prev
			scores[player] += marble.value
			marble.prev.next = marble.next
			marble.next.prev = marble.prev
			cur = marble.next
		} else {
			marble := &playedMarble{nextValue, cur.next, cur.next.next}
			cur.next.next.prev = marble
			cur.next.next = marble
			cur = marble
		}
	}
	return max(scores)
}

func max(values []int) int {
	m := 0
	for _, v := range values {
		if m < v {
			m = v
		}
	}
	return m
}
