package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	notes := parseNotes(io.ReadLines())
	sum := 0
	nums := map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}
	for _, n := range notes {
		var s [10][]byte
		var l5 [][]byte
		var l6 [][]byte
		m := map[byte]byte{}

		for _, signal := range n.signals {
			switch len(signal) {
			case 2:
				s[1] = signal
			case 3:
				s[7] = signal
			case 4:
				s[4] = signal
			case 5:
				l5 = append(l5, signal)
			case 6:
				l6 = append(l6, signal)
			case 7:
				s[8] = signal
			}
		}

		// 'a'
		m['a'] = subtract(s[7], s[1])[0]

		// 'g'
		half9 := append(s[4], m['a'])
		m['g'], s[9] = detectWithOneDifference(l6, half9)

		// 'e'
		m['e'] = subtract(s[8], s[9])[0]

		// 'f'
		s[2] = detectWithTwoDifference(l5, []byte{m['a'], m['e'], m['g']})
		m['f'] = subtract(s[7], s[2])[0]

		// 'c'
		m['c'] = subtract(s[7], []byte{m['a'], m['f']})[0]

		// 'd'
		m['d'] = subtract(s[2], []byte{m['a'], m['c'], m['e'], m['g']})[0]

		// 'b'
		m['b'] = subtract(s[8], []byte{m['a'], m['c'], m['d'], m['e'], m['f'], m['g']})[0]

		rm := map[byte]byte{}
		for wrong, correct := range m {
			rm[correct] = wrong
		}

		output := 0
		for _, wrong := range n.output {
			var correct sortableBytes
			for _, w := range wrong {
				correct = append(correct, rm[w])
			}
			sort.Sort(correct)
			output *= 10
			output += nums[string(correct)]
		}

		// for wrong, correct := range rm {
		// 	fmt.Printf("%c => %c, ", wrong, correct)
		// }
		// fmt.Println(output)
		sum += output
	}
	fmt.Println(sum)
}

func detectWithOneDifference(minuends [][]byte, subtrahend []byte) (difference byte, detected []byte) {
	for _, minuend := range minuends {
		d := subtract(minuend, subtrahend)
		if len(d) == 1 {
			difference = d[0]
			detected = minuend
			break
		}
	}
	return
}

func detectWithTwoDifference(minuends [][]byte, subtrahend []byte) (detected []byte) {
	for _, minuend := range minuends {
		d := subtract(minuend, subtrahend)
		if len(d) == 2 {
			detected = minuend
			break
		}
	}
	return
}

func subtract(minuend []byte, subtrahend []byte) (difference []byte) {
	for _, m := range minuend {
		keep := true
		for _, s := range subtrahend {
			if m == s {
				keep = false
				break
			}
		}
		if keep {
			difference = append(difference, m)
		}
	}
	return
}

func parseNotes(lines []string) (notes []*note) {
	for _, line := range lines {
		raw := strings.Split(line, " | ")
		notes = append(notes, &note{
			signals: strings2Bytes(strings.Split(raw[0], " ")),
			output:  strings2Bytes(strings.Split(raw[1], " ")),
		})
	}
	return
}

func strings2Bytes(strings []string) (bytes [][]byte) {
	for _, s := range strings {
		bytes = append(bytes, []byte(s))
	}
	return
}

type note struct {
	signals [][]byte
	output  [][]byte
}

type sortableBytes []byte

var _ sort.Interface = (sortableBytes)(nil)

func (s sortableBytes) Len() int {
	return len(s)
}

func (s sortableBytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortableBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
