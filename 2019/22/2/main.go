package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/22/spacecards"
)

func main() {
	// seen := map[int]bool{}
	input, _ := ioutil.ReadAll(os.Stdin)
	idx := 2020
	i := 0
	// for !seen[idx] {
	for {
		// seen[idx] = true
		fmt.Print(i, idx)
		idx = spacecards.ReverseLookup(string(input), 119315717514047, idx)
		fmt.Println(":", idx)
		i++
		// if len(seen)%100000 == 0 {
		// 	fmt.Println(idx, len(seen))
		// }
		if i > 10 {
			break
		}
	}
}
