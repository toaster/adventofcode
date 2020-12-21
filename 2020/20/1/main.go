package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type image struct {
	id                           int
	pixels                       [][]bool
	top, left, right, bottom     int
	rtop, rleft, rright, rbottom int
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	var images []*image
	idsAndImages := strings.Split(strings.Trim(string(input), "\n"), "\n\n")
	for _, idAndImage := range idsAndImages {
		images = append(images, parseImage(idAndImage))
	}
	edgeRefCounts := map[int]int{}
	edgesToImages := map[int][]*image{}
	edgesToFlippedImages := map[int][]*image{}
	outerEdgeCountByImage := map[*image]int{}
	outerEdgeCountByFlippedImage := map[*image]int{}
	for _, img := range images {
		edgeRefCounts[img.top]++
		edgeRefCounts[img.left]++
		edgeRefCounts[img.right]++
		edgeRefCounts[img.bottom]++
		edgeRefCounts[img.rtop]++
		edgeRefCounts[img.rleft]++
		edgeRefCounts[img.rright]++
		edgeRefCounts[img.rbottom]++
		edgesToImages[img.top] = append(edgesToImages[img.top], img)
		edgesToImages[img.left] = append(edgesToImages[img.left], img)
		edgesToImages[img.right] = append(edgesToImages[img.right], img)
		edgesToImages[img.bottom] = append(edgesToImages[img.bottom], img)
		edgesToFlippedImages[img.rtop] = append(edgesToFlippedImages[img.rtop], img)
		edgesToFlippedImages[img.rleft] = append(edgesToFlippedImages[img.rleft], img)
		edgesToFlippedImages[img.rright] = append(edgesToFlippedImages[img.rright], img)
		edgesToFlippedImages[img.rbottom] = append(edgesToFlippedImages[img.rbottom], img)
	}
	for e, c := range edgeRefCounts {
		if c == 1 {
			if len(edgesToImages[e]) == 0 {
				outerEdgeCountByFlippedImage[edgesToFlippedImages[e][0]]++
			} else {
				outerEdgeCountByImage[edgesToImages[e][0]]++
			}
		}
	}
	cornerIDs := map[int]bool{}
	for _, img := range images {
		if outerEdgeCountByImage[img] == 2 || outerEdgeCountByFlippedImage[img] == 2 {
			cornerIDs[img.id] = true
		}
	}
	product := 1
	for id, _ := range cornerIDs {
		product *= id
	}
	fmt.Println("product:", product)
}

func parseImage(input string) *image {
	idAndLines := strings.Split(input, "\n")
	id, _ := strconv.Atoi(idAndLines[0][5:9])
	img := &image{id: id, pixels: make([][]bool, 10)}
	for y, line := range idAndLines[1:] {
		img.pixels[y] = make([]bool, 10)
		for x, pix := range line {
			img.pixels[y][x] = pix == '#'
		}
	}
	for i := 0; i < 10; i++ {
		if img.pixels[0][9-i] {
			img.top += 1 << i
		}
		if img.pixels[0][i] {
			img.rtop += 1 << i
		}
		if img.pixels[i][0] {
			img.left += 1 << i
		}
		if img.pixels[9-i][0] {
			img.rleft += 1 << i
		}
		if img.pixels[9-i][9] {
			img.right += 1 << i
		}
		if img.pixels[i][9] {
			img.rright += 1 << i
		}
		if img.pixels[9][9-i] {
			img.rbottom += 1 << i
		}
		if img.pixels[9][i] {
			img.bottom += 1 << i
		}
	}
	return img
}
