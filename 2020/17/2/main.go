package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type point struct {
	x int
	y int
	z int
	w int
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	pocketDimension := map[point]bool{}
	min := point{0, 0, 0, 0}
	max := point{0, 0, 0, 0}
	for y, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		for x, state := range line {
			if state == '#' {
				pocketDimension[point{x, y, 0, 0}] = true
			}
		}
		max.x = len(line) - 1
		max.y = y
	}

	for i := 0; i < 6; i++ {
		pocketDimension = cycle(pocketDimension, &min, &max)
	}
	fmt.Println("active:", len(pocketDimension))
}

func cycle(dim map[point]bool, min, max *point) map[point]bool {
	sx := min.x - 1
	ex := max.x + 1
	sy := min.y - 1
	ey := max.y + 1
	sz := min.z - 1
	ez := max.z + 1
	sw := min.w - 1
	ew := max.w + 1
	newDim := map[point]bool{}
	for w := sw; w <= ew; w++ {
		for z := sz; z <= ez; z++ {
			for y := sy; y <= ey; y++ {
				for x := sx; x <= ex; x++ {
					p := point{x, y, z, w}
					c := countActiveNeighbours(dim, p)
					if (dim[p] && (c == 2 || c == 3)) || (!dim[p] && c == 3) {
						newDim[p] = true
					}
					if newDim[p] {
						if p.x < min.x {
							min.x = p.x
						}
						if p.y < min.y {
							min.y = p.y
						}
						if p.z < min.z {
							min.z = p.z
						}
						if p.w < min.w {
							min.w = p.w
						}
						if p.x > max.x {
							max.x = p.x
						}
						if p.y > max.y {
							max.y = p.y
						}
						if p.z > max.z {
							max.z = p.z
						}
						if p.w > max.w {
							max.w = p.w
						}
					}
				}
			}
		}
	}
	return newDim
}

func countActiveNeighbours(dim map[point]bool, p point) int {
	active := 0
	for w := p.w - 1; w <= p.w+1; w++ {
		for z := p.z - 1; z <= p.z+1; z++ {
			for y := p.y - 1; y <= p.y+1; y++ {
				for x := p.x - 1; x <= p.x+1; x++ {
					n := point{x, y, z, w}
					if n == p {
						continue
					}
					if dim[n] {
						active++
					}
				}
			}
		}
	}
	return active
}

// func printDim(dim map[point]bool, min, max point) {
// 	for z := min.z; z <= max.z; z++ {
// 		fmt.Printf("z=%d; %v; %v\n", z, min, max)
// 		for y := min.y; y <= max.y; y++ {
// 			for x := min.x; x <= max.x; x++ {
// 				if dim[point{x, y, z}] {
// 					fmt.Print("#")
// 				} else {
// 					fmt.Print(".")
// 				}
// 			}
// 			fmt.Println()
// 		}
// 		fmt.Println()
// 	}
// }
