package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	lines := io.ReadLines()
	var wins [][]int
	var nums [][]int
	var cards []int
	for _, line := range lines {
		raw := strings.Split(strings.Split(line, ": ")[1], " | ")
		wins = append(wins, io.ParseInts(raw[0], " "))
		nums = append(nums, io.ParseInts(raw[1], " "))
		cards = append(cards, 1)
	}
	sum := 0
	for i := 0; i < len(wins); i++ {
		common := math.Intersection(nums[i], wins[i])
		count := len(common)
		fmt.Printf("Card %d wins %d (%d - %d)\n", i+1, count, i+1+1, i+1+1+count)
		for j := i + 1; j < len(cards) && j < i+1+count; j++ {
			cards[j] += cards[i]
		}
		sum += cards[i]
	}
	fmt.Printf("cards: %d\n", cards)
	fmt.Println(sum)
}
