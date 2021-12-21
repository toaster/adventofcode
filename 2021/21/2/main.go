package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

var advanceCounts = map[int]int{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

func main() {
	lines := io.ReadLines()
	pos1 := parseLine(lines[0])
	pos2 := parseLine(lines[1])
	fieldCount := 10
	winScore := 21
	gw1, gl1 := findGroups(pos1, 0, 0, fieldCount, winScore)
	gw2, gl2 := findGroups(pos2, 0, 0, fieldCount, winScore)
	fmt.Println(gw1, gl1, "\n", gw2, gl2)
	wonRounds1 := 0
	for rounds, c := range gw1 {
		wonRounds1 += c * gl2[rounds-1]
	}
	wonRounds2 := 0
	for rounds, c := range gw2 {
		wonRounds2 += c * gl1[rounds]
	}
	fmt.Println(wonRounds1, wonRounds2)
}

func findGroups(pos, score, round, fieldCount, winScore int) (map[int]int, map[int]int) {
	winGroups := map[int]int{}
	loseGroups := map[int]int{}
	newRound := round + 1
	for i := 3; i < 10; i++ {
		newPos := ((pos + i - 1) % fieldCount) + 1
		newScore := score + newPos
		ac := advanceCounts[i]
		if newScore < winScore {
			wg, lg := findGroups(newPos, newScore, newRound, fieldCount, winScore)
			for s, c := range wg {
				winGroups[s] += c * ac
			}
			for s, c := range lg {
				loseGroups[s] += c * ac
			}
			loseGroups[newRound] += ac
		} else {
			winGroups[newRound] += ac
		}
	}
	return winGroups, loseGroups
}

func parseLine(line string) int {
	value, err := strconv.Atoi(strings.Split(line, ": ")[1])
	io.ReportError("", err)
	return value
}
