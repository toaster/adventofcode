package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
	"github.com/toaster/advent_of_code/internal/math"
)

func main() {
	risks := map[math.Point2D]int{}
	lines := io.ReadLines()
	height := len(lines)
	width := len(lines[0])
	nodes := map[math.Point2D]*math.Node{}
	for y, line := range lines {
		for x, c := range line {
			p := math.Point2D{X: x, Y: y}
			risks[p] = int(c) - '0'
			nodes[p] = &math.Node{}
		}
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p := math.Point2D{X: x, Y: y}
			n := nodes[p]
			if x > 0 {
				addEdge(nodes, math.Point2D{X: x - 1, Y: y}, n, risks)
			}
			if x < width-1 {
				addEdge(nodes, math.Point2D{X: x + 1, Y: y}, n, risks)
			}
			if y > 0 {
				addEdge(nodes, math.Point2D{X: x, Y: y - 1}, n, risks)
			}
			if y < height-1 {
				addEdge(nodes, math.Point2D{X: x, Y: y + 1}, n, risks)
			}
		}
	}
	start := nodes[math.Point2D{X: 0, Y: 0}]
	end := nodes[math.Point2D{X: width - 1, Y: height - 1}]

	fmt.Println(math.ComputeLengthOfShortestPath(start, end))

	// pathPoints := map[math.Point2D]bool{end.Data.(math.Point2D): true}
	// for _, node := range end.Path {
	// 	pathPoints[node.Data.(math.Point2D)] = true
	// }
	// length := 0
	// for y := 0; y < height; y++ {
	// 	for x := 0; x < width; x++ {
	// 		p := math.Point2D{X: x, Y: y}
	// 		// n := nodes[p]
	// 		// fmt.Printf("%02d ", n.Distance)
	// 		if pathPoints[p] {
	// 			fmt.Print(risks[p])
	// 			length += risks[p]
	// 		} else {
	// 			fmt.Print(" ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println(length, risks[start.Data.(math.Point2D)], length-risks[start.Data.(math.Point2D)])
}

func addEdge(nodes map[math.Point2D]*math.Node, point math.Point2D, node *math.Node, risks map[math.Point2D]int) {
	targetNode := nodes[point]
	node.Edges = append(node.Edges, &math.Edge{
		Weight: risks[point],
		A:      node,
		B:      targetNode,
	})
}
