package spacecards

import (
	"strconv"
	"strings"
)

// Deck is a deck of space cards of a certain size.
type Deck struct {
	Cards []int
}

// NewDeck creates a new deck of space cards in factory order of given size.
func NewDeck(size int) *Deck {
	c := make([]int, size)
	for i := 0; i < size; i++ {
		c[i] = i
	}
	return &Deck{Cards: c}
}

// Shuffle applies the shuffle steps specified in input on the deck.
func (d *Deck) Shuffle(input string) {
	for _, cmd := range strings.Split(strings.TrimSpace(input), "\n") {
		if strings.HasPrefix(cmd, "cut") {
			i, _ := strconv.Atoi(cmd[4:])
			d.cut(i)
		}
		if strings.HasPrefix(cmd, "deal with increment") {
			i, _ := strconv.Atoi(cmd[20:])
			d.dealWithIncrement(i)
		}
		if cmd == "deal into new stack" {
			d.dealIntoNewStack()
		}
	}
}

func ReverseLookup(input string, count, idx int) int {
	cmds := strings.Split(strings.TrimSpace(input), "\n")
	for i := len(cmds) - 1; i >= 0; i-- {
		cmd := cmds[i]
		if strings.HasPrefix(cmd, "cut") {
			n, _ := strconv.Atoi(cmd[4:])
			if n < 1 {
				n = count + n
			}
			idx = idx + n
			if idx >= count-n {
				idx -= count
			}
		}
		if strings.HasPrefix(cmd, "deal with increment") {
			n, _ := strconv.Atoi(cmd[20:])
			idx = ((n-(idx%n))*count + idx) / n
		}
		if cmd == "deal into new stack" {
			idx = count - idx - 1
		}
	}
	return idx
}

func (d *Deck) cut(i int) {
	n := make([]int, 0, len(d.Cards))
	if i < 0 {
		i = len(d.Cards) + i
	}
	n = append(n, d.Cards[i:]...)
	n = append(n, d.Cards[0:i]...)
	d.Cards = n
}

func (d *Deck) dealWithIncrement(i int) {
	n := make([]int, len(d.Cards))
	p := 0
	for _, card := range d.Cards {
		n[p] = card
		p += i
		p %= len(d.Cards)
	}
	d.Cards = n
}

func (d *Deck) dealIntoNewStack() {
	n := make([]int, 0, len(d.Cards))
	for i := len(d.Cards) - 1; i >= 0; i-- {
		n = append(n, d.Cards[i])
	}
	d.Cards = n
}
