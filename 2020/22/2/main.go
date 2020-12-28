package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var results = map[string]int{}
var hits = 0

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "failed reading standard input:", err)
		os.Exit(1)
	}

	var deckA []int
	var deckB []int
	for i, block := range strings.Split(strings.Trim(string(input), "\n"), "\n\n") {
		for _, line := range strings.Split(block, "\n") {
			if strings.HasPrefix(line, "Player") {
				continue
			}

			num, _ := strconv.Atoi(line)
			if i == 0 {
				deckA = append(deckA, num)
			} else {
				deckB = append(deckB, num)
			}
		}
	}
	score, _ := playGame([]int{1}, deckA, deckB)
	fmt.Println("winner score:", score)
}

func playGame(games []int, deckA []int, deckB []int) (int, bool) {
	gameKey := join(deckA, ",") + "-" + join(deckB, ",")
	if results[gameKey] > 0 {
		hits++
		// fmt.Println("winner", results[gameKey], "---", len(results), "---", hits)
		return 0, results[gameKey] == 1
	}
	reverseGameKey := join(deckB, ",") + "-" + join(deckA, ",")
	var seenDecks []string
	subgame := 0
	game := join(games, ".")
	fmt.Printf("=== Game %s ===\n\n", game)
	round := 0
	for len(deckA) > 0 && len(deckB) > 0 {
		round++
		fmt.Printf("-- Round %d (Game %s) --\n", round, game)
		fmt.Println("Player 1’s deck:", join(deckA, ", "))
		fmt.Println("Player 2’s deck:", join(deckB, ", "))
		a := deckA[0]
		b := deckB[0]
		fmt.Println("Player 1 plays:", a)
		fmt.Println("Player 2 plays:", b)
		deckA = deckA[1:]
		deckB = deckB[1:]
		var aWins bool
		if a <= len(deckA) && b <= len(deckB) {
			fmt.Print("Playing a sub-game to determine the winner…\n\n")
			subgame++
			subDeckA := make([]int, a)
			subDeckB := make([]int, b)
			copy(subDeckA, deckA)
			copy(subDeckB, deckB)
			_, aWins = playGame(append(games, subgame), subDeckA, subDeckB)
			fmt.Printf("…anyway, back to game %s.\n", game)
		} else {
			aWins = a > b
		}
		winner := 1
		if aWins {
			deckA = append(deckA, a, b)
		} else {
			winner = 2
			deckB = append(deckB, b, a)
		}
		sA := join(deckA, ",")
		sB := join(deckB, ",")
		fmt.Printf("Player %d wins round %d of game %s!\n\n", winner, round, game)
		for _, deck := range seenDecks {
			if deck == sA || deck == sB {
				fmt.Printf("The winner of game %s is player 1! (Rescue)\n\n", game)
				results[gameKey] = 1
				results[reverseGameKey] = 2
				return calculateScore(deckA), true
			}
		}
		seenDecks = append(seenDecks, sA, sB)
	}
	winner := 1
	winnerDeck := deckA
	if len(deckA) == 0 {
		winnerDeck = deckB
		winner = 2
	}
	results[gameKey] = winner
	results[reverseGameKey] = winner%2 + 1
	// fmt.Printf("The winner of game %s is player %d!\n\n", game, winner)
	score := calculateScore(winnerDeck)
	return score, len(deckA) != 0
}

func calculateScore(deck []int) (score int) {
	for i, num := range deck {
		score += num * (len(deck) - i)
	}
	return
}

func join(nums []int, sep string) string {
	var s []string
	for _, num := range nums {
		s = append(s, strconv.Itoa(num))
	}
	return strings.Join(s, sep)
}
