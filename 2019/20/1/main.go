package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/20/pluto"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	m := pluto.Parse(string(input))
	fmt.Println(m.ShortestPath())
}
