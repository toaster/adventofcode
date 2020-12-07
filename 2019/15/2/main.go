package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/15/repairdroid"
	"github.com/toaster/advent_of_code/2019/icc"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	p := icc.Parse(string(input))
	r := repairdroid.Explore(p)
	r.PrintMap()
	fmt.Println("time:", r.FillWithOxygen())
}
