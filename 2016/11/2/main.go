package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	const floorCount = 4
	objects := map[string]int{}
	var isotopes []string
	for floor, line := range io.ReadLines() {
		words := strings.Split(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						line,
						".",
						"",
					),
					",",
					"",
				),
				"and ",
				"",
			),
			" ",
		)
		for i := 4; i < len(words); i += 3 {
			if words[i] != "a" {
				break
			}

			isotope := strings.ToUpper(words[i+1][:1])
			if !slices.Contains(isotopes, isotope) {
				isotopes = append(isotopes, isotope)
			}
			id := isotope
			if words[i+2] == "microchip" {
				id += "M"
			} else {
				id += "G"
			}
			objects[id] = floor + 1
		}
	}
	for _, isotope := range "ED" {
		isotopes = append(isotopes, string(isotope))
		objects[string(isotope)+"G"] = 1
		objects[string(isotope)+"M"] = 1
	}
	slices.Sort(isotopes)
	elevatorLocation := 1
	steps := 0

	printMap := func() {
		fmt.Println("after", steps, "steps")
		for floor := floorCount; floor > 0; floor-- {
			fmt.Printf("F%d ", floor)
			if floor == elevatorLocation {
				fmt.Print("E ")
			} else {
				fmt.Print(". ")
			}
			for _, isotope := range isotopes {
				id := isotope + "G"
				if objects[id] == floor {
					fmt.Printf("%s ", id)
				} else {
					fmt.Print(".  ")
				}
				id = isotope + "M"
				if objects[id] == floor {
					fmt.Printf("%s ", id)
				} else {
					fmt.Print(".  ")
				}
			}
			fmt.Println()
		}
	}
	printMap()

	transport := func(distance int, ids []string) {
		if len(ids) < 1 || len(ids) > 2 || math.AbsInt(distance) != 1 {
			panic("elevator mis-usage")
		}

		elevatorLocation += distance
		for _, id := range ids {
			objects[id] = elevatorLocation
		}
		steps++
		printMap()
	}
	up := func(ids ...string) {
		transport(1, ids)
	}
	down := func(ids ...string) {
		transport(-1, ids)
	}

	// “SG” is the main elevator boy
	elevatorBoy := "SG"
	// “TG” and “PG” pair to prepare the set-up for the elevator boy.
	// First, get the first floor clear except for the chip of the elevator boy.
	up("PG", "PM")
	up("PG", "PM")
	down("PG")
	down("PG")
	up("DG", "DM")
	down("TG")
	up("PG", "TG")
	down("PG")
	up("EG", "EM")
	down("TG")
	up("PG", "TG")
	up("PG", "TG")
	up("PG", "PM")
	down("PG")
	up("PG", "TG")
	down("TG")
	down("TG")
	down("TG")
	// Then, get the elevator boy and let it transfer the T pair to the top.
	up("TG", elevatorBoy)
	up("TG", elevatorBoy)
	up("TG", elevatorBoy)
	down("TG")
	up("TG", "TM")
	// Now, the elevator boy is at the top and the third floor is clean which is the start position for “fetchFromSecondFloor”.
	fetchFromSecondFloor := func(elevatorBoy, isotope string) {
		down(elevatorBoy)
		down(elevatorBoy)
		generator := isotope + "G"
		chip := isotope + "M"
		up(generator, chip)
		down(generator)
		up(elevatorBoy, generator)
		up(elevatorBoy, generator)
		down(generator)
		up(generator, chip)
	}
	fetchFromSecondFloor(elevatorBoy, "C")
	fetchFromSecondFloor(elevatorBoy, "R")
	fetchFromSecondFloor(elevatorBoy, "E")
	fetchFromSecondFloor(elevatorBoy, "D")
	down(elevatorBoy)
	down(elevatorBoy)
	down(elevatorBoy)
	up(elevatorBoy, "SM")
	up(elevatorBoy, "SM")
	up(elevatorBoy, "SM")
}
