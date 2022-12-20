package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	var nodes []*node
	var cur *node
	var zero *node
	for _, line := range io.ReadLines() {
		cur = &node{
			prev:  cur,
			value: io.ParseInt(line),
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

	mix(nodes)
	n := nodeAt(zero, 1000)
	v1 := n.value
	n = nodeAt(n, 1000)
	v2 := n.value
	n = nodeAt(n, 1000)
	v3 := n.value
	fmt.Println(v1 + v2 + v3)
}

func mix(nodes []*node) {
	for _, n := range nodes {
		if n.value < 0 {
			for i := 0; i > n.value; i-- {
				n.prev.next = n.next
				n.next.prev = n.prev
				n.next = n.prev
				n.prev = n.next.prev
				n.next.prev = n
				n.prev.next = n
			}
		} else {
			for i := 0; i < n.value; i++ {
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
