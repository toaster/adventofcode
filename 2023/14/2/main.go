package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2023/14/reflectordish"
	"github.com/toaster/advent_of_code/internal/io"
)

const (
	ball rockType = 'O'
	cube rockType = '#'
	none rockType = '.'
)

type rockType rune

func main() {
	p := reflectordish.ParseControlPlatform(io.ReadLines())
	periodLength := 0
	secondPeriodStart := 0
	loads := []int{p.NorthLoad()}
	const cycles = 1000000000
	{
		assumedPeriodLength := 0
		verifiedPeriodSteps := 0
		lastSeen := map[int]int{}
		for i := 1; i <= cycles; i++ {
			p.SpinCycle()
			load := p.NorthLoad()
			loads = append(loads, load)
			if assumedPeriodLength > 0 {
				if loads[i-assumedPeriodLength] == load {
					verifiedPeriodSteps++
					if verifiedPeriodSteps == assumedPeriodLength {
						periodLength = assumedPeriodLength
						secondPeriodStart = i
						break
					}
				} else {
					assumedPeriodLength = 0
					verifiedPeriodSteps = 0
				}
			}
			if assumedPeriodLength == 0 {
				if lastSeen[load] > 0 {
					assumedPeriodLength = i - lastSeen[load]
					if assumedPeriodLength < 3 {
						assumedPeriodLength = 0
					}
				}
			}
			lastSeen[load] = i
		}
	}
	periods := cycles / periodLength
	index := periods*periodLength + secondPeriodStart - cycles - 1
	fmt.Println(loads[index])
}
