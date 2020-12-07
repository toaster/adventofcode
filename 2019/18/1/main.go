package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/18/dungeon"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	m := dungeon.Parse(string(input))
	fmt.Println(m.MinimalStepsToCollectAllKeys())
}
