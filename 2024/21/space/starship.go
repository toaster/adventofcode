package space

import (
	"slices"
	"strings"

	"github.com/toaster/advent_of_code/internal/math"
)

// ShortestSequence computes the length of the shortest sequence of keystrokes on a directional
// keypad which controls a chain of robots as long as the given indirections which eventually types
// on a numerical keypad the requested code.
// https://adventofcode.com/2024/day/21
func ShortestSequence(code string, indirections int) int {
	moves := map[string]bool{}
	numSequences := newNumericKeypad().Enter(code)
	for _, sequence := range numSequences {
		for _, move := range sequence {
			moves[move] = true
		}
	}

	cache := map[string]*cacheEntry{}
	for indirection := 0; indirection < indirections; indirection++ {
		next := map[string]bool{}
		for move := range moves {
			if cache[move] != nil {
				continue
			}

			sequences := newDirectionalKeypad().Enter(move)
			for _, sequence := range sequences {
				for _, m := range sequence {
					next[m] = true
				}
			}
			cache[move] = &cacheEntry{sequences: sequences, indirectionCounts: map[int][]int{}}
		}
		moves = next
	}

	minCount := 0
	for _, sequence := range numSequences {
		count := 0
		for _, move := range sequence {
			count += computeCount(move, cache, indirections-1)
		}
		if minCount == 0 || minCount > count {
			minCount = count
		}
	}
	return minCount
}

func computeCount(move string, cache map[string]*cacheEntry, indirection int) int {
	entry := cache[move]
	minCount := 0
	if entry.indirectionCounts[indirection] == nil {
		entry.indirectionCounts[indirection] = make([]int, len(entry.sequences))
	}
	for i, moves := range entry.sequences {
		if entry.indirectionCounts[indirection][i] == 0 {
			if indirection == 0 {
				l := 0
				for _, m := range moves {
					l += len(m)
				}
				entry.indirectionCounts[indirection][i] = l
			} else {
				count := 0
				for _, m := range moves {
					count += computeCount(m, cache, indirection-1)
				}
				entry.indirectionCounts[indirection][i] = count
			}
		}
		if minCount == 0 || minCount > entry.indirectionCounts[indirection][i] {
			minCount = entry.indirectionCounts[indirection][i]
		}
	}
	return minCount
}

func newDirectionalKeypad() *keypad {
	return newKeypadWithKeys(map[rune]math.Point2D{
		'^': {X: 1},
		'>': {X: 2, Y: 1},
		'v': {X: 1, Y: 1},
		'<': {Y: 1},
		'A': {X: 2},
	})
}

func newKeypadWithKeys(keys map[rune]math.Point2D) *keypad {
	positions := map[math.Point2D]bool{}
	for _, pos := range keys {
		positions[pos] = true
	}
	return &keypad{
		cur:       keys['A'],
		keys:      keys,
		positions: positions,
	}
}

func newNumericKeypad() *keypad {
	return newKeypadWithKeys(map[rune]math.Point2D{
		'0': {X: 1, Y: 3},
		'1': {Y: 2},
		'2': {X: 1, Y: 2},
		'3': {X: 2, Y: 2},
		'4': {Y: 1},
		'5': {X: 1, Y: 1},
		'6': {X: 2, Y: 1},
		'7': {},
		'8': {X: 1},
		'9': {X: 2},
		'A': {X: 2, Y: 3},
	})
}

type cacheEntry struct {
	sequences         [][]string
	indirectionCounts map[int][]int
}

type keypad struct {
	cur       math.Point2D
	keys      map[rune]math.Point2D
	positions map[math.Point2D]bool
}

func (p *keypad) Enter(input string) [][]string {
	moves := [][]string{nil}
	for _, key := range input {
		possibilities := p.press(key)
		c := len(moves)
		for i := 0; i < c; i++ {
			if len(possibilities) > 1 {
				moves = append(moves, append(slices.Clone(moves[i]), possibilities[1]))
			}
			moves[i] = append(moves[i], possibilities[0])
		}
	}
	return moves
}

func (p *keypad) press(key rune) []string {
	target := p.keys[key]
	if target == p.cur {
		return []string{"A"}
	}

	xOffset := target.X - p.cur.X
	yOffset := target.Y - p.cur.Y
	var moves []string
	if _, ok := p.positions[p.cur.AddXY(xOffset, 0)]; xOffset != 0 && ok {
		keys := &strings.Builder{}
		p.writeMovement(xOffset, '<', '>', keys)
		p.writeMovement(yOffset, '^', 'v', keys)
		keys.WriteRune('A')
		moves = append(moves, keys.String())
	}
	if _, ok := p.positions[p.cur.AddXY(0, yOffset)]; yOffset != 0 && ok {
		keys := &strings.Builder{}
		p.writeMovement(yOffset, '^', 'v', keys)
		p.writeMovement(xOffset, '<', '>', keys)
		keys.WriteRune('A')
		moves = append(moves, keys.String())
	}
	p.cur = target
	return moves
}

func (p *keypad) writeMovement(offset int, decr, incr rune, keys *strings.Builder) {
	var k rune
	if offset < 0 {
		k = decr
		offset = -offset
	} else {
		k = incr
	}
	for i := 0; i < offset; i++ {
		keys.WriteRune(k)
	}
}
