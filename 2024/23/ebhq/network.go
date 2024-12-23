package ebhq

import (
	"maps"
	"slices"
	"strings"
)

// NewNetwork creates a new Network
func NewNetwork() *Network {
	return &Network{nodes: map[string]*node{}}
}

// Network represents a network in the Easter Bunny Headquarter (https://adventofcode.com/2024/day/23).
type Network struct {
	nodes map[string]*node
}

// AddConnection adds a connection to the Network.
func (n *Network) AddConnection(idA, idB string) {
	if _, ok := n.nodes[idA]; !ok {
		n.nodes[idA] = &node{id: idA, connections: map[string]*node{}}
	}
	if _, ok := n.nodes[idB]; !ok {
		n.nodes[idB] = &node{id: idB, connections: map[string]*node{}}
	}

	n.nodes[idA].connect(n.nodes[idB])
	n.nodes[idB].connect(n.nodes[idA])
}

// ComputeLANPartyPassword computes the password to THE LAN party, i.e., to the largest network group.
func (n *Network) ComputeLANPartyPassword() string {
	p := n.nodes
	r := map[string]*node{}
	x := map[string]*node{}
	ids := n.bronKerbosch(r, p, x)
	slices.Sort(ids)
	return strings.Join(ids, ",")
}

// CountGroupsOfThree counts all groups of three interconnected nodes in the Network.
func (n *Network) CountGroupsOfThree() int {
	threeNodesSets := map[string]bool{}
	for id, nd := range n.nodes {
		if !strings.HasPrefix(id, "t") {
			continue
		}

		for i, cidA := range nd.connectedIDs {
			for _, cidB := range nd.connectedIDs[i+1:] {
				if _, ok := nd.connections[cidA].connections[cidB]; !ok {
					continue
				}

				set := []string{id, cidA, cidB}
				slices.Sort(set)
				setID := strings.Join(set, ",")
				if !threeNodesSets[setID] {
					threeNodesSets[setID] = true
				}
			}
		}
	}
	return len(threeNodesSets)
}

func (n *Network) bronKerbosch(r, p, x map[string]*node) []string {
	if len(p) == 0 && len(x) == 0 {
		var ids []string
		for id := range r {
			ids = append(ids, id)
		}
		return ids
	}

	var maxClique []string
	originalP := maps.Clone(p)
	for id, nd := range originalP {
		newR := maps.Clone(r)
		newR[id] = nd
		newP := map[string]*node{}
		newX := map[string]*node{}
		for _, connectedID := range nd.connectedIDs {
			if tgt, ok := p[connectedID]; ok {
				newP[connectedID] = tgt
			}
			if tgt, ok := x[connectedID]; ok {
				newX[connectedID] = tgt
			}
		}
		candidate := n.bronKerbosch(newR, newP, newX)
		if len(candidate) > len(maxClique) {
			maxClique = candidate
		}
		x[id] = p[id]
		delete(p, id)
	}
	return maxClique
}

type node struct {
	connections  map[string]*node
	connectedIDs []string
	id           string
}

func (n *node) connect(other *node) {
	if _, ok := n.connections[other.id]; ok {
		return
	}
	n.connections[other.id] = other
	n.connectedIDs = append(n.connectedIDs, other.id)
	slices.Sort(n.connectedIDs)
}
