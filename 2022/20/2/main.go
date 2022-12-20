package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	key := 811589153
	var nodes []*node
	var cur *node
	var zero *node
	lines := io.ReadLines()
	for _, line := range lines {
		cur = &node{
			prev:  cur,
			value: io.ParseInt(line) * key,
		}
		if cur.value == 0 {
			zero = cur
		}
		nodes = append(nodes, cur)
	}
	nodes[0].prev = cur
	cur.next = nodes[0]
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].next = nodes[i+1]
	}

	for i := 0; i < 10; i++ {
		mix(nodes)
	}
	n := nodeAt(zero, 1000)
	v1 := n.value
	n = nodeAt(n, 1000)
	v2 := n.value
	n = nodeAt(n, 1000)
	v3 := n.value
	fmt.Println(v1 + v2 + v3)
}

func mix(nodes []*node) {
	maxSteps := len(nodes) - 1
	for _, n := range nodes {
		v := n.value % maxSteps
		if v < 0 {
			for i := 0; i > v; i-- {
				n.prev.next = n.next
				n.next.prev = n.prev
				n.next = n.prev
				n.prev = n.next.prev
				n.next.prev = n
				n.prev.next = n
			}
		} else {
			for i := 0; i < v; i++ {
				n.prev.next = n.next
				n.next.prev = n.prev
				n.prev = n.next
				n.next = n.prev.next
				n.next.prev = n
				n.prev.next = n
			}
		}
	}
}

func nodeAt(n *node, steps int) *node {
	for i := 0; i < steps; i++ {
		n = n.next
	}
	return n
}

type node struct {
	next  *node
	prev  *node
	value int
}
