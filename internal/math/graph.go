package math

import "fmt"

// Edge is an edge of a weighted undirected graph connecting two Nodes.
type Edge struct {
	Weight int
	A      *Node
	B      *Node
}

// Node is a node of a weighted undirected graph.
type Node struct {
	Edges []*Edge
}

type depthFirstPath struct {
	length  int
	steps   []*Node
	visited map[*Node]bool
}

// ComputeLengthOfShortestPath efficiently computes the length of the shortest path between two nodes
// in a weighted undirected graph.
func ComputeLengthOfShortestPath(start *Node, end *Node) int {
	distances := map[*Node]int{start: 1}
	candidates := map[*Node]int{}
	cur := start
	for cur != end {
		var next *Node
		for node, costs := range candidates {
			if next == nil || costs < candidates[next] {
				next = node
			}
		}
		for _, edge := range cur.Edges {
			var node *Node
			if edge.A == cur {
				node = edge.B
			} else {
				node = edge.A
			}
			if distances[node] != 0 {
				continue
			}

			dist := distances[cur] + edge.Weight
			if candidates[node] == 0 || candidates[node] > dist {
				candidates[node] = dist
			}
			if next == nil || candidates[node] < candidates[next] {
				next = node
			}
		}
		distances[next] = candidates[next]
		delete(candidates, next)
		cur = next
	}
	return distances[end] - 1
}

// ComputeShortestPath computes the shortest path between two nodes in a weighted undirected graph.
func ComputeShortestPath(start *Node, end *Node) ([]*Node, int) {
	shortest := shortestDepthFirstPath(&depthFirstPath{steps: []*Node{start}, visited: map[*Node]bool{start: true}}, end)
	return nil, shortest.length
}

// CountPossibleDestinations is a broad-first implementation to search for all reachable nodes within a given distance in an unweighted graph.
func CountPossibleDestinations[T comparable](adjacents func(T) []T, start T, maxDistance int) int {
	visited := map[T]bool{start: true}
	next := []T{start}
	distance := 0
	for {
		cur := next
		next = nil
		for _, pos := range cur {
			visited[pos] = true
			for _, node := range adjacents(pos) {
				if !visited[node] {
					next = append(next, node)
				}
			}
		}
		distance++
		if distance > maxDistance {
			return len(visited)
		}
	}
}

// FindShortestDistance is a broad-first implementation to search for the shortest path in an unweighted graph.
func FindShortestDistance[T comparable](adjacents func(T) []T, start, end T) int {
	return FindShortestWeightedDistance(adjacents, func(_, _ T) int { return 1 }, func(p T) bool { return p == end }, start)
}

// FindShortestWeightedDistance is a broad-first implementation to search for the shortest path in a weighted graph.
func FindShortestWeightedDistance[T comparable](adjacents func(T) []T, weight func(T, T) int, isEnd func(T) bool, start T) int {
	visited := map[T]bool{start: true}
	candidates := map[int][]T{0: {start}}
	distance := 0
	maxDistance := 0
	for {
		if distance > maxDistance {
			fmt.Printf("unexpected end\n")
			return -1
		}
		for _, pos := range candidates[distance] {
			if isEnd(pos) {
				return distance
			}
			for _, node := range adjacents(pos) {
				if !visited[node] {
					d := distance + weight(pos, node)
					if d > maxDistance {
						maxDistance = d
					}
					candidates[d] = append(candidates[d], node)
				}
				visited[node] = true
			}
		}
		distance++
	}
}

// PossibleDestinationsForExactDistance is a broad-first implementation to search for all reachable nodes for a given distance in an unweighted graph.
// Nodes which are within shorter distance are not included unless there is a path to them with exactly the given distance.
func PossibleDestinationsForExactDistance[T comparable](adjacents func(T) []T, start T, targetDistance int) map[T]bool {
	next := map[T]bool{start: true}
	distance := 0
	for {
		cur := next
		if distance == targetDistance {
			return cur
		}

		next = map[T]bool{}
		for pos := range cur {
			for _, node := range adjacents(pos) {
				next[node] = true
			}
		}
		distance++
		fmt.Printf("\r%d", distance)
	}
}

func shortestDepthFirstPath(startPath *depthFirstPath, end *Node) (shortestPath *depthFirstPath) {
	var paths []*depthFirstPath
	for _, edge := range startPath.steps[len(startPath.steps)-1].Edges {
		newPath := copyDepthFirstPath(startPath)
		// TODO: error: dest node is B if edge.A == start node (startPath.steps[len(startPath.steps)-1])
		// TODO: continue if dest node is already visited
		var node *Node
		if startPath.visited[edge.A] {
			node = edge.B
		} else {
			node = edge.A
		}
		newPath.steps = append(newPath.steps, node)
		newPath.visited[node] = true
		newPath.length += edge.Weight
		if node == end {
			paths = append(paths, newPath)
		} else {
			paths = append(paths, shortestDepthFirstPath(newPath, end))
		}
	}
	for _, path := range paths {
		if shortestPath == nil || path.length < shortestPath.length {
			shortestPath = path
		}
	}
	return
}

func copyDepthFirstPath(path *depthFirstPath) *depthFirstPath {
	newSteps := make([]*Node, len(path.steps))
	copy(newSteps, path.steps)
	newVisited := make(map[*Node]bool, len(path.visited))
	for node, visited := range path.visited {
		newVisited[node] = visited
	}
	return &depthFirstPath{
		length:  path.length,
		steps:   newSteps,
		visited: newVisited,
	}
}
