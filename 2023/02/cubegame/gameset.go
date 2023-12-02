package cubegame

import (
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

// ParseRecords parses an input of records and returns the result GameSet.
func ParseRecords(lines []string) (games GameSet) {
	for _, line := range lines {
		components := strings.Split(line, ": ")
		id := io.ParseInt(components[0][5:])
		g := &game{id: id}
		for _, rawRecord := range strings.Split(components[1], "; ") {
			r := &record{}
			for _, rawSet := range strings.Split(rawRecord, ", ") {
				valueAndColor := strings.Split(rawSet, " ")
				value := io.ParseInt(valueAndColor[0])
				switch valueAndColor[1] {
				case "red":
					r.red = value
				case "green":
					r.green = value
				case "blue":
					r.blue = value
				}
			}
			g.records = append(g.records, r)
		}
		games = append(games, g)
	}
	return
}

type game struct {
	id      int
	records []*record
}

// GameSet describes a set of cube games played.
type GameSet []*game

// SumOfPossibleGames returns the sum of the Game IDs of the game set which are possible to play with a given bag content.
func (s GameSet) SumOfPossibleGames(red, green, blue int) (sum int) {
	for _, g := range s {
		possible := true
		for _, r := range g.records {
			if r.red > red || r.green > green || r.blue > blue {
				possible = false
				break
			}
		}
		if possible {
			sum += g.id
		}
	}
	return
}

// SumOfPowersOfMinimumSets returns the sum of the powers of minimal cube sets of all games in the set.
func (s GameSet) SumOfPowersOfMinimumSets() (sum int) {
	for _, g := range s {
		red, green, blue := 0, 0, 0
		for _, r := range g.records {
			red = max(red, r.red)
			green = max(green, r.green)
			blue = max(blue, r.blue)
		}
		sum += red * green * blue
	}
	return
}

type record struct {
	red   int
	green int
	blue  int
}
