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

	images := map[int]*image{}
	idsAndImages := strings.Split(strings.Trim(string(input), "\n"), "\n\n")
	for _, idAndImage := range idsAndImages {
		img := parseImage(idAndImage)
		images[img.id] = img
	}
	edgeRefCounts := map[int]int{}
	imagesByEdges := map[int][]*image{}
	flippedImagesByEdges := map[int][]*image{}
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
		imagesByEdges[img.top] = append(imagesByEdges[img.top], img)
		imagesByEdges[img.left] = append(imagesByEdges[img.left], img)
		imagesByEdges[img.right] = append(imagesByEdges[img.right], img)
		imagesByEdges[img.bottom] = append(imagesByEdges[img.bottom], img)
		flippedImagesByEdges[img.rtop] = append(flippedImagesByEdges[img.rtop], img)
		flippedImagesByEdges[img.rleft] = append(flippedImagesByEdges[img.rleft], img)
		flippedImagesByEdges[img.rright] = append(flippedImagesByEdges[img.rright], img)
		flippedImagesByEdges[img.rbottom] = append(flippedImagesByEdges[img.rbottom], img)
	}
	for e, c := range edgeRefCounts {
		if c == 1 {
			if len(imagesByEdges[e]) == 0 {
				outerEdgeCountByFlippedImage[flippedImagesByEdges[e][0]]++
			} else {
				outerEdgeCountByImage[imagesByEdges[e][0]]++
			}
		}
	}
	cornerIDs := map[int]bool{}
	edgeIDs := map[int]bool{}
	innerIDs := map[int]bool{}
	for _, img := range images {
		if outerEdgeCountByImage[img] == 2 || outerEdgeCountByFlippedImage[img] == 2 {
			cornerIDs[img.id] = true
		} else if outerEdgeCountByImage[img] == 1 || outerEdgeCountByFlippedImage[img] == 1 {
			edgeIDs[img.id] = true
		} else {
			innerIDs[img.id] = true
		}

	}
	dim := len(edgeIDs)/4 + 2
	tileIDs := make([][]int, dim)
	for y := range tileIDs {
		tileIDs[y] = make([]int, dim)
	}
	for id := range cornerIDs {
		img := images[id]
		for edgeRefCounts[img.top] != 1 || edgeRefCounts[img.left] != 1 {
			rotate(img)
		}
		tileIDs[0][0] = id
		break
	}
	completeRow(dim, 0, tileIDs, images, imagesByEdges, flippedImagesByEdges)
	for y := 1; y < dim; y++ {
		{
			prev := tileIDs[y-1][0]
			var candidates []int
			edge := images[prev].bottom
			for _, img := range imagesByEdges[edge] {
				if img.id == prev {
					continue
				}
				flip(img)
				candidates = append(candidates, img.id)
			}
			for _, img := range flippedImagesByEdges[edge] {
				if img.id == prev {
					continue
				}
				candidates = append(candidates, img.id)
			}
			if len(candidates) != 1 {
				panic("oopsi")
			}
			img := images[candidates[0]]
			for img.rtop != edge {
				rotate(img)
			}
			tileIDs[y][0] = img.id
		}
		completeRow(dim, y, tileIDs, images, imagesByEdges, flippedImagesByEdges)
	}
	img := compose(images, tileIDs)
	c, r := countMonstersAndRoughness(img)
	fmt.Println("c:", c, "r: ", r)
	rotate(img)
	c, r = countMonstersAndRoughness(img)
	fmt.Println("c:", c, "r: ", r)
	rotate(img)
	c, r = countMonstersAndRoughness(img)
	fmt.Println("c:", c, "r: ", r)
	rotate(img)
	c, r = countMonstersAndRoughness(img)
	fmt.Println("c:", c, "r: ", r)
	flip(img)
	c, r = countMonstersAndRoughness(img)
	fmt.Println("c:", c, "r: ", r)
	rotate(img)
	c, r = countMonstersAndRoughness(img)
	fmt.Println("c:", c, "r: ", r)
	rotate(img)
	c, r = countMonstersAndRoughness(img)
	fmt.Println("c:", c, "r: ", r)
	rotate(img)
	c, r = countMonstersAndRoughness(img)
	fmt.Println("c:", c, "r: ", r)
}

func completeRow(dim int, y int, tileIDs [][]int, images map[int]*image, imagesByEdges map[int][]*image, flippedImagesByEdges map[int][]*image) {
	for x := 1; x < dim; x++ {
		prev := tileIDs[y][x-1]
		var candidates []int
		edge := images[prev].right
		for _, img := range imagesByEdges[edge] {
			if img.id == prev {
				continue
			}
			flip(img)
			candidates = append(candidates, img.id)
		}
		for _, img := range flippedImagesByEdges[edge] {
			if img.id == prev {
				continue
			}
			candidates = append(candidates, img.id)
		}
		if len(candidates) != 1 {
			panic("oopsi")
		}
		img := images[candidates[0]]
		for img.rleft != edge {
			rotate(img)
		}
		tileIDs[y][x] = img.id
	}
}

func compose(images map[int]*image, tileIDs [][]int) *image {
	img := &image{id: -1}
	for _, row := range tileIDs {
		for y := 1; y < 9; y++ {
			pixRow := make([]bool, len(row)*8)
			img.pixels = append(img.pixels, pixRow)
			for ix, id := range row {
				tile := images[id]
				for x := 1; x < 9; x++ {
					pixRow[ix*8+x-1] = tile.pixels[y][x]
				}
			}
		}
	}
	// for _, row := range img.pixels {
	// 	for _, p := range row {
	// 		if p {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	return img
}

func computeEdges(img *image) {
	img.top = 0
	img.left = 0
	img.right = 0
	img.bottom = 0
	img.rtop = 0
	img.rleft = 0
	img.rright = 0
	img.rbottom = 0
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
}

func countMonstersAndRoughness(img *image) (int, int) {
	height := len(img.pixels)
	width := len(img.pixels[0])
	monsterCount := 0
	roughCount := 0
	fmt.Println(width, "x", height)
	for y, row := range img.pixels {
		for x, p := range row {
			if x < width-19 && y < height-2 {
				if img.pixels[y][x+18] &&
					img.pixels[y+1][x+19] && img.pixels[y+1][x+18] && img.pixels[y+1][x+17] &&
					img.pixels[y+1][x+12] && img.pixels[y+1][x+11] &&
					img.pixels[y+1][x+6] && img.pixels[y+1][x+5] &&
					img.pixels[y+1][x] &&
					img.pixels[y+2][x+16] &&
					img.pixels[y+2][x+13] && img.pixels[y+2][x+10] &&
					img.pixels[y+2][x+7] && img.pixels[y+2][x+4] &&
					img.pixels[y+2][x+1] {
					monsterCount++
				}
			}
			if p {
				roughCount++
			}

		}
	}
	return monsterCount, roughCount - monsterCount*15
}

func flip(img *image) {
	height := len(img.pixels)
	newPixels := make([][]bool, height)
	for y := 0; y < height; y++ {
		newPixels[y] = img.pixels[height-1-y]
	}
	img.pixels = newPixels
	computeEdges(img)
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
	computeEdges(img)
	// printTile(img)
	return img
}

func printTile(img *image) {
	fmt.Println("ID:", img.id)
	fmt.Println("    ", img.top, "-", img.rtop)
	fmt.Printf("% 4d ", img.left)
	for _, p := range img.pixels[0] {
		if p {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("", img.right)
	fmt.Printf("% 4d ", img.rleft)
	for _, p := range img.pixels[1] {
		if p {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("", img.rright)
	for y := 2; y < 10; y++ {
		fmt.Print("     ")
		for _, p := range img.pixels[y] {
			if p {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("    ", img.bottom, "-", img.rbottom)
	fmt.Println()
}

func rotate(img *image) {
	height := len(img.pixels)
	width := len(img.pixels[0])
	newPixels := make([][]bool, width)
	for y := 0; y < width; y++ {
		newPixels[y] = make([]bool, height)
		for x := 0; x < height; x++ {
			newPixels[y][x] = img.pixels[x][width-1-y]
		}
	}
	img.pixels = newPixels
	computeEdges(img)
}
