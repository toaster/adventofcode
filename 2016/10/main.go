package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	bots := map[int]*robot{}
	outputs := map[int]chan int{}
	for _, line := range io.ReadLines() {
		words := strings.Split(line, " ")
		if words[0] == "value" {
			i := io.ParseInt(words[5])
			bot := getBot(bots, i)
			input := bot.createInput()
			value := io.ParseInt(words[1])
			go func() {
				input <- value
			}()
		} else {
			i := io.ParseInt(words[1])
			bot := getBot(bots, i)
			lowI := io.ParseInt(words[6])
			highI := io.ParseInt(words[11])
			var low, high chan int
			if words[5] == "bot" {
				otherBot := getBot(bots, lowI)
				low = otherBot.createInput()
			} else {
				outputs[lowI] = make(chan int, 2)
				low = outputs[lowI]
			}
			if words[10] == "bot" {
				otherBot := getBot(bots, highI)
				high = otherBot.createInput()
			} else {
				outputs[highI] = make(chan int)
				high = outputs[highI]
			}
			bot.outputs = []chan int{low, high}
		}
	}
	for _, bot := range bots {
		bot.run()
	}
	time.Sleep(1 * time.Second)
	for i, bot := range bots {
		if bot.compared[0] == 17 && bot.compared[1] == 61 {
			fmt.Println("Part 1:", i, "=>", bot.compared)
		}
	}
	product := 1
	for i := 0; i < 3; i++ {
		v := <-outputs[i]
		product *= v
	}
	fmt.Println("Part 2:", product)
}

func getBot(bots map[int]*robot, i int) *robot {
	if bots[i] == nil {
		bots[i] = &robot{}
	}
	return bots[i]
}

type robot struct {
	inputs   []chan int
	outputs  []chan int
	compared []int
}

func (r *robot) createInput() chan int {
	input := make(chan int)
	r.inputs = append(r.inputs, input)
	return input
}

func (r *robot) run() {
	go func() {
		var values []int
		for _, input := range r.inputs {
			values = append(values, <-input)
		}
		slices.Sort(values)
		r.compared = values
		for i, output := range r.outputs {
			output <- values[i]
		}
	}()
}
