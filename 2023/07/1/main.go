package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

const (
	highCard handType = iota
	onePair
	twoPairs
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

var labelValue = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func main() {
	lines := io.ReadLines()
	var hands []*hand
	for _, line := range lines {
		components := strings.Split(line, " ")
		h := &hand{}
		copy(h.cards[:], components[0])
		h.bidAmount = io.ParseInt(components[1])
		h.analyze()
		hands = append(hands, h)
	}
	slices.SortStableFunc(hands, func(a, b *hand) int {
		return a.compare(b)
	})
	totalWinning := 0
	for i, h := range hands {
		totalWinning += h.bidAmount * (i + 1)
	}
	fmt.Println(totalWinning)
}

type hand struct {
	cards     [5]byte
	typ       handType
	bidAmount int
}

func (h *hand) analyze() {
	counts := map[byte]int{}
	for _, card := range h.cards {
		counts[card]++
	}
	for _, count := range counts {
		switch count {
		case 5:
			h.typ = fiveOfAKind
		case 4:
			h.typ = fourOfAKind
		case 3:
			if h.typ == onePair {
				h.typ = fullHouse
			} else {
				h.typ = threeOfAKind
			}
		case 2:
			if h.typ == threeOfAKind {
				h.typ = fullHouse
			} else if h.typ == onePair {
				h.typ = twoPairs
			} else {
				h.typ = onePair
			}
		}
	}
}

func (h *hand) compare(other *hand) int {
	if h.typ < other.typ {
		return -1
	}
	if h.typ > other.typ {
		return 1
	}

	for i := 0; i < len(h.cards); i++ {
		myLabelValue := labelValue[h.cards[i]]
		otherLabelValue := labelValue[other.cards[i]]
		if myLabelValue == 0 {
			panic(string(h.cards[:]))
		}
		if otherLabelValue == 0 {
			panic(string(other.cards[:]))
		}
		if myLabelValue < otherLabelValue {
			return -1
		}
		if myLabelValue > otherLabelValue {
			return 1
		}
	}

	return 0
}

func (h *hand) describe() string {
	t := "high card"
	switch h.typ {
	case fiveOfAKind:
		t = "five of a kind"
	case fourOfAKind:
		t = "four of a kind"
	case fullHouse:
		t = "full house"
	case threeOfAKind:
		t = "three of a kind"
	case twoPairs:
		t = "two pairs"
	case onePair:
		t = "one pair"
	}
	return fmt.Sprintf("cards: %s type: %s bid: %d", string(h.cards[:]), t, h.bidAmount)
}

type handType int
