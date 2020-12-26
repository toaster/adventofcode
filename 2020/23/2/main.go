package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type cup struct {
	label int
	next  *cup
	prev  *cup
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	var current, c *cup
	cups := map[int]*cup{}
	for _, n := range strings.Split(strings.Trim(string(input), "\n"), "") {
		num, _ := strconv.Atoi(n)
		nc := &cup{label: num}
		if c == nil {
			current = nc
		} else {
			c.next = nc
			nc.prev = c
		}
		cups[num] = nc
		c = nc
	}
	cupCount := 1000000
	for i := 10; i <= cupCount; i++ {
		nc := &cup{label: i}
		c.next = nc
		nc.prev = c
		cups[i] = nc
		c = nc
	}
	c.next = current
	current.prev = c
	moveCount := cupCount * 10
	for move := 0; move < moveCount; move++ {
		pick := current.next
		current.next = pick.next.next.next
		pick.next.next.next.prev = current
		pick.prev = nil
		pick.next.next.next = nil
		destLabel := current.label - 1
		if destLabel < 1 {
			destLabel = cupCount
		}
		for {
			done := true
			for p := pick; p != nil; p = p.next {
				if p.label == destLabel {
					destLabel--
					if destLabel < 1 {
						destLabel = cupCount
					}
					done = false
					break
				}
			}
			if done {
				break
			}
		}
		dest := cups[destLabel]
		pick.next.next.next = dest.next
		dest.next.prev = pick.next.next
		dest.next = pick
		pick.prev = dest
		current = current.next
	}
	result := cups[1].next.label * cups[1].next.next.label
	fmt.Println("result:", result)
}
