package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type ground byte

const (
	lumberyard ground = '#'
	open       ground = '.'
	trees      ground = '|'
)

type pos struct {
	x, y int
}

type area struct {
	m             map[pos]ground
	width, height int
}

func main() {
	inputFile := os.Args[1]
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	minutes := 10
	if len(os.Args) > 2 {
		minutes, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
	}

	lines := strings.Split(strings.TrimSuffix(string(b), "\n"), "\n")
	a := area{m: map[pos]ground{}, width: len(lines[0]), height: len(lines)}
	for y, l := range lines {
		for x, g := range []ground(l) {
			a.m[pos{x, y}] = g
		}
	}

	history := []area{a}

	printMap(a)
	periodStart := 0
	periodLength := 0
	for i := 1; i <= minutes && periodStart == 0; i++ {
		a = evolve(a)
		history = append(history, a)
		// fmt.Println("After", i, "minutes:")
		// printMap(a)
		for j := i - 1; j > i-50 && j >= 0; j-- {
			if equal(a, history[j]) {
				fmt.Println("period from", j, "through", i, "length:", i-j)
				periodStart = j
				periodLength = i - j
				break
			}
		}
	}
	if periodStart > 0 {
		index := (minutes-periodStart)%periodLength + periodStart
		a = history[index]
	}
	fmt.Println("After", minutes, "minutes:")
	printMap(a)
	wood := 0
	lumber := 0
	for y := 0; y < a.height; y++ {
		for x := 0; x < a.width; x++ {
			switch (a.m[pos{x, y}]) {
			case lumberyard:
				lumber++
			case trees:
				wood++
			}
		}
	}
	fmt.Println(wood * lumber)
}

func printMap(a area) {
	for y := 0; y < a.height; y++ {
		for x := 0; x < a.width; x++ {
			fmt.Print(string(a.m[pos{x, y}]))
		}
		fmt.Println("")
	}
}

func evolve(a area) area {
	b := area{m: map[pos]ground{}, width: a.width, height: a.height}
	for y := 0; y < a.height; y++ {
		for x := 0; x < a.width; x++ {
			p := pos{x, y}
			ad := adjacents(a, p)
			switch a.m[p] {
			case open:
				if count(ad, trees) >= 3 {
					b.m[p] = trees
				} else {
					b.m[p] = a.m[p]
				}
			case lumberyard:
				if count(ad, trees) > 0 && count(ad, lumberyard) > 0 {
					b.m[p] = a.m[p]
				} else {
					b.m[p] = open
				}
			case trees:
				if count(ad, lumberyard) >= 3 {
					b.m[p] = lumberyard
				} else {
					b.m[p] = a.m[p]
				}
			default:
				b.m[p] = a.m[p]
			}
		}
	}
	return b
}

func adjacents(a area, p pos) []ground {
	r := []ground{}
	if p.x > 0 {
		r = append(r, a.m[pos{p.x - 1, p.y}])
		if p.y > 0 {
			r = append(r, a.m[pos{p.x - 1, p.y - 1}])
		}
		if p.y < a.height-1 {
			r = append(r, a.m[pos{p.x - 1, p.y + 1}])
		}
	}
	if p.y > 0 {
		r = append(r, a.m[pos{p.x, p.y - 1}])
	}
	if p.y < a.height-1 {
		r = append(r, a.m[pos{p.x, p.y + 1}])
	}
	if p.x < a.width-1 {
		r = append(r, a.m[pos{p.x + 1, p.y}])
		if p.y > 0 {
			r = append(r, a.m[pos{p.x + 1, p.y - 1}])
		}
		if p.y < a.height-1 {
			r = append(r, a.m[pos{p.x + 1, p.y + 1}])
		}
	}
	return r
}

func count(gs []ground, g ground) int {
	c := 0
	for _, x := range gs {
		if x == g {
			c++
		}
	}
	return c
}

func equal(a, b area) bool {
	if a.width != b.width || a.height != b.height {
		return false
	}
	for k, v := range a.m {
		if b.m[k] != v {
			return false
		}
	}
	return true
}
