package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/24/eris"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	m := eris.New(string(input))
	fmt.Println(m.SimulateUntilRepeat())
}
