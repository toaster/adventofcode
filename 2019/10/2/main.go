package main

import (
	"fmt"
	"io"
	"os"

	"github.com/toaster/advent_of_code/2019/space"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	m := space.ParseMap(string(input))
	fmt.Println(m.VaporizeAsteroids(math.Point2D{X: 25, Y: 31})[199])
}
