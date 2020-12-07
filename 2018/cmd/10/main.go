package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type pixel struct {
	x, y, dx, dy int
}

func main() {
	inputRegex := regexp.MustCompile("position=< *([-0-9]+), *([-0-9]+)> velocity=< *([-0-9]+), *([-0-9]+)>")
	inputFile := os.Args[1]
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	pixels := make([]*pixel, len(lines))
	seconds := 0
	for i, line := range lines {
		if line == "" {
			continue
		}
		matches := inputRegex.FindStringSubmatch(line)
		pixels[i] = &pixel{}
		pixels[i].x, err = strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		pixels[i].y, err = strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		pixels[i].dx, err = strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}
		pixels[i].dy, err = strconv.Atoi(matches[4])
		if err != nil {
			panic(err)
		}
	}
	done := false
	delta := 0
	for !done {
		seconds += 1
		var minY, maxY int
		var minX, maxX int
		for i, pixel := range pixels {
			if pixel == nil {
				continue
			}
			pixel.x += pixel.dx
			pixel.y += pixel.dy
			if i == 0 || minX > pixel.x {
				minX = pixel.x
			}
			if i == 0 || maxX < pixel.x {
				maxX = pixel.x
			}
			if i == 0 || minY > pixel.y {
				minY = pixel.y
			}
			if i == 0 || maxY < pixel.y {
				maxY = pixel.y
			}
		}
		if delta != 0 && maxY-minY > delta {
			for i, pixel := range pixels {
				if pixel == nil {
					continue
				}
				pixel.x -= pixel.dx
				pixel.y -= pixel.dy
				if i == 0 || minX > pixel.x {
					minX = pixel.x
				}
				if i == 0 || maxX < pixel.x {
					maxX = pixel.x
				}
				if i == 0 || minY > pixel.y {
					minY = pixel.y
				}
				if i == 0 || maxY < pixel.y {
					maxY = pixel.y
				}
			}
			done = true
			w := maxX - minX + 1
			h := maxY - minY + 1
			d := make([]byte, w*h)
			for i, _ := range d {
				d[i] = ' '
			}
			for _, pixel := range pixels {
				if pixel == nil {
					continue
				}
				d[(pixel.y-minY)*w+pixel.x-minX] = '#'
			}
			for i := 0; i < h; i++ {
				fmt.Println(string(d[i*w : (i+1)*w]))
			}
			fmt.Println(seconds-1, "seconds")
		}
		delta = maxY - minY
	}
}
