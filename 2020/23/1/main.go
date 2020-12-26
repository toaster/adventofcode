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

	var cups []int
	for _, n := range strings.Split(strings.Trim(string(input), "\n"), "") {
		num, _ := strconv.Atoi(n)
		cups = append(cups, num)
	}
	current := 0
	for move := 0; move < 100; move++ {
		var pick []int
		pickStart := (current + 1) % 9
		fmt.Println("-- move", move+1, "--")
		fmt.Print("cups:")
		for i, l := range cups {
			if i == current {
				fmt.Printf(" (%d)", l)
			} else {
				fmt.Printf(" %d", l)
			}
		}
		fmt.Println()
		if pickStart == 8 {
			pick = append(pick, cups[8], cups[0], cups[1])
			cups = cups[2:8]
			current -= 2
		} else if pickStart == 7 {
			pick = append(pick, cups[7], cups[8], cups[0])
			cups = cups[1:7]
			current--
		} else {
			pick = append(pick, cups[pickStart:pickStart+3]...)
			if pickStart == 0 {
				cups = cups[3:]
				current -= 3
			} else {
				cups = append(cups[0:pickStart], cups[pickStart+3:]...)
			}
		}
		fmt.Print("pick up:")
		for _, l := range pick {
			fmt.Printf(" %d", l)
		}
		fmt.Println()
		destLabel := cups[current] - 1
		if destLabel < 1 {
			destLabel = 9
		}
		for {
			done := true
			for _, l := range pick {
				if l == destLabel {
					destLabel--
					if destLabel < 1 {
						destLabel = 9
					}
					done = false
					break
				}
			}
			if done {
				break
			}
		}
		fmt.Println("destination:", destLabel)
		for i, l := range cups {
			if l == destLabel {
				if i == 8 {
					cups = append(pick, cups...)
					current += 3
				} else if i == 7 {
					cups = append(pick[1:], cups...)
					cups = append(cups, pick[0])
					current++
				} else if i == 6 {
					cups = append(pick[2:], cups...)
					cups = append(cups, pick[0], pick[1])
					current += 2
				} else {
					cups = append(cups[0:i+1], append(pick, cups[i+1:]...)...)
					if i < current {
						current += 3
					}
				}
			}
		}
		current = (current + 1) % 9
	}
	order := ""
	collect := false
	for _, l := range cups {
		if l == 1 {
			collect = true
			continue
		}
		if !collect {
			continue
		}
		order = order + strconv.Itoa(l)
	}
	for _, l := range cups {
		if l == 1 {
			break
		}
		order = order + strconv.Itoa(l)
	}
	fmt.Println("order:", order)
}

func calculateScore(deck []int) (score int) {
	for i, num := range deck {
		score += num * (len(deck) - i)
	}
	return
}
