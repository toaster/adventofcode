package math

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
