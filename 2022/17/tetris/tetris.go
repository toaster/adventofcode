package tetris

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/math"
	"github.com/toaster/advent_of_code/internal/util"
)

// Piece represents a piece in a game of Tetris.
type Piece struct {
	Height int
	Shape  map[math.Point2D]bool
	Width  int
}

type board struct {
	Width             int
	OccupiedPositions map[math.Point2D]bool
}

// Play plays a game of Tetris on a board with the given width using the given pieces and sidewards directions
// until the specified account of pieces has settled.
func Play(width, pieceCount int, pieces []*Piece, jets []rune) int {
	settledCount := 0
	highestPoint := 0
	pieceIndex := 0
	jetIndex := 0
	b := &board{Width: width, OccupiedPositions: map[math.Point2D]bool{}}
	for settledCount < pieceCount {
		piece := pieces[pieceIndex]
		pieceIndex = nextIndex(pieces, pieceIndex)
		pos := math.Point2D{X: 3, Y: highestPoint + 4}
		settled := false
		for !settled {
			xDelta := 0
			switch jets[jetIndex] {
			case '>':
				xDelta = 1
			case '<':
				xDelta = -1
			}
			newPos := pos.AddXY(xDelta, 0)
			jetIndex = nextIndex(jets, jetIndex)
			if !collides(b, piece, newPos) {
				pos = newPos
			}

			newPos = pos.AddXY(0, -1)
			if collides(b, piece, newPos) {
				settled = true
			} else {
				pos = newPos
			}
		}
		settledCount++
		for x := 0; x < piece.Width; x++ {
			for y := 0; y < piece.Height; y++ {
				p := math.Point2D{X: x, Y: y}
				if piece.Shape[p] {
					b.OccupiedPositions[pos.Add(p)] = true
				}
			}
		}
		highestPoint = math.MaxInt(highestPoint, pos.Y+piece.Height-1)
		// printBoard(b, highestPoint)
		// fmt.Println(highestPoint)
	}
	printBoard(b, highestPoint)
	return highestPoint
}

// PlayWithShortcut plays a game of Tetris just like Play, but it tries to determine a pattern to
// reduce the actual computation time (i.e. it should be able to handle ridiculous high piece counts).
func PlayWithShortcut(width, pieceCount int, pieces []*Piece, jets []rune) int {
	settledCount := 0
	highestPoint := 0
	pieceIndex := 0
	jetIndex := 0
	var occurrences []math.Point2D
	var highestPoints []int
	period := 0
	periodCount := 0
	b := &board{Width: width, OccupiedPositions: map[math.Point2D]bool{}}
	for settledCount < pieceCount {
		highestPoints = append(highestPoints, highestPoint)
		occurrence := math.Point2D{X: pieceIndex, Y: jetIndex}
		if p := checkOccurrencePattern(occurrences, occurrence); p > 0 {
			if period == 0 {
				period = p
			}
			if period == p {
				periodCount++
			} else {
				// period changed -> not stable yet, reset.
				period = 0
				periodCount = 0
			}
			if periodCount == period {
				fmt.Println("detected confident period", period, "at", settledCount)
			}
			if periodCount >= period {
				fmt.Println("potential growth", highestPoints[len(highestPoints)-2]-highestPoints[len(highestPoints)-period-2], highestPoint-highestPoints[len(highestPoints)-period-1])
				remainingPieces := pieceCount - settledCount
				periodMatchesRemaining := remainingPieces%period == 0
				fmt.Println("check if", remainingPieces, "matches period", periodMatchesRemaining)
				if periodMatchesRemaining {
					growth := highestPoint - highestPoints[len(highestPoints)-period-1]
					remainingPeriods := remainingPieces / period
					return highestPoint + remainingPeriods*growth
				}
			}
		}
		occurrences = append(occurrences, occurrence)
		piece := pieces[pieceIndex]
		pieceIndex = nextIndex(pieces, pieceIndex)
		pos := math.Point2D{X: 3, Y: highestPoint + 4}
		settled := false
		for !settled {
			xDelta := 0
			switch jets[jetIndex] {
			case '>':
				xDelta = 1
			case '<':
				xDelta = -1
			}
			newPos := pos.AddXY(xDelta, 0)
			jetIndex = nextIndex(jets, jetIndex)
			if !collides(b, piece, newPos) {
				pos = newPos
			}

			newPos = pos.AddXY(0, -1)
			if collides(b, piece, newPos) {
				settled = true
			} else {
				pos = newPos
			}
		}
		settledCount++
		for x := 0; x < piece.Width; x++ {
			for y := 0; y < piece.Height; y++ {
				p := math.Point2D{X: x, Y: y}
				if piece.Shape[p] {
					b.OccupiedPositions[pos.Add(p)] = true
				}
			}
		}
		highestPoint = math.MaxInt(highestPoint, pos.Y+piece.Height-1)
		// printBoard(b, highestPoint)
		// fmt.Println(highestPoint)
	}
	// printBoard(b, highestPoint)
	return highestPoint
}

func checkOccurrencePattern(occurrences []math.Point2D, occurrence math.Point2D) int {
	for i := len(occurrences) - 1; i >= 0; i-- {
		if occurrences[i] == occurrence {
			p := len(occurrences) - i
			if i >= p && util.SlicesEqual(occurrences[i:], occurrences[i-p:i]) {
				return p
			}
		}
	}
	return 0
}

func collides(b *board, piece *Piece, pos math.Point2D) bool {
	if pos.X <= 0 || pos.X+piece.Width >= b.Width+2 {
		// hit the wall
		return true
	}
	if pos.Y <= 0 {
		// hit the ground
		return true
	}

	for x := 0; x < piece.Width; x++ {
		for y := 0; y < piece.Height; y++ {
			p := math.Point2D{X: x, Y: y}
			if piece.Shape[p] && b.OccupiedPositions[pos.Add(p)] {
				return true
			}
		}
	}

	return false
}

func nextIndex[T any](slice []T, index int) int {
	index++
	if index == len(slice) {
		return 0
	}
	return index
}

func printBoard(b *board, top int) {
	for y := top; y >= 0; y-- {
		for x := 0; x < b.Width+2; x++ {
			if x == 0 || x == b.Width+1 {
				if y == 0 {
					fmt.Print("+")
				} else {
					fmt.Print("|")
				}
				continue
			}
			if y == 0 {
				fmt.Print("-")
				continue
			}
			if b.OccupiedPositions[math.Point2D{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
