package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/space"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	m := space.Parse(string(input))
	fmt.Println(m.VaporizeAsteroids(space.Point{X: 25, Y: 31})[199])
}
