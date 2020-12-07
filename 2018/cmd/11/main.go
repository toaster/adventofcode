package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	w := 300
	h := 300
	sStart := 3
	sEnd := 3
	serial, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 3 {
		sStart, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
		sEnd, err = strconv.Atoi(os.Args[3])
		if err != nil {
			panic(err)
		}
	}
	grid := make([]int, w*h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			rackID := x + 10
			grid[y*w+x] = (((rackID*y+serial)*rackID)/100)%10 - 5
		}
	}
	maxX := 0
	maxY := 0
	maxPower := 0
	maxS := 0
	for s := sStart; s < sEnd+1; s++ {
		fmt.Println(s, ",", maxPower)
		for y := 0; y < h+1-s; y++ {
			for x := 0; x < w+1-s; x++ {
				power := 0
				for i := 0; i < s; i++ {
					for j := 0; j < s; j++ {
						power += grid[(y+i)*w+x+j]
					}
				}
				if maxPower < power {
					maxPower = power
					maxX = x
					maxY = y
					maxS = s
				}
			}
		}
	}
	fmt.Println(maxX, ",", maxY, ",", maxS, ":", maxPower)
	for y := maxY - 1; y < maxY+4; y++ {
		for x := maxX - 1; x < maxX+4; x++ {
			fmt.Printf("% 3d", grid[y*w+x])
		}
		fmt.Println("")
	}
}
