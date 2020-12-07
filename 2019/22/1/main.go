package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/toaster/advent_of_code/2019/22/spacecards"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	d := spacecards.NewDeck(10007)
	d.Shuffle(string(input))
	for i, card := range d.Cards {
		if card == 2019 {
			fmt.Println(i)
			break
		}
	}
}
