package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/06/1/orbit"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println(orbit.Count(string(input)))
}
