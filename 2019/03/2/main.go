package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/toaster/advent_of_code/2019/03/2/grid"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	lines := strings.Split(string(input), "\n")
	g := map[grid.Pos]int{}
	grid.Mark(g, lines[0])
	fmt.Println(grid.Search(g, lines[1]))
}
