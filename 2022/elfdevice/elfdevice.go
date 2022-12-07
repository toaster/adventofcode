package elfdevice

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

// FSNode represents a node in the filesystem of the elf device.
type FSNode struct {
	Entries map[string]*FSNode
	IsDir   bool
	Name    string
	Parent  *FSNode
	Size    int
}

// ParseShellSession parses the output of a filesystem crawling shell session.
func ParseShellSession() *FSNode {
	root := &FSNode{Entries: map[string]*FSNode{}, IsDir: true, Name: "/"}
	cur := root
	for _, line := range io.ReadLines() {
		components := strings.Split(line, " ")
		if components[0] == "$" {
			if components[1] == "cd" {
				name := components[2]
				switch name {
				case "/":
					cur = root
				case "..":
					cur = cur.Parent
				default:
					cur = cur.Entries[name]
				}
			}
		} else {
			entry := &FSNode{}
			entry.Name = components[1]
			entry.Parent = cur
			if components[0] == "dir" {
				entry.IsDir = true
				entry.Entries = map[string]*FSNode{}
			} else {
				entry.Size = io.ParseInt(components[0])
				for d := cur; d != nil; d = d.Parent {
					d.Size += entry.Size
				}
			}
			cur.Entries[entry.Name] = entry
			continue
		}
	}
	return root
}

// PrintNode prints the whole subtree of an FSNode.
func PrintNode(node *FSNode, level int) {
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += " "
	}
	fmt.Printf("%s- %s ", prefix, node.Name)
	if node.IsDir {
		fmt.Printf("(dir, size=%d)\n", node.Size)
		for _, child := range node.Entries {
			PrintNode(child, level+2)
		}
	} else {
		fmt.Printf("(file, size=%d)\n", node.Size)
	}
}

// SmallestDirWithSizeAtLeast returns the smallest directory in the given tree which size is at
// least minSize or `nil` if there is no such directory.
func SmallestDirWithSizeAtLeast(dir *FSNode, minSize int) *FSNode {
	if dir.Size < minSize {
		return nil
	}
	cur := dir
	for _, node := range dir.Entries {
		if node.IsDir {
			if x := SmallestDirWithSizeAtLeast(node, minSize); x != nil && x.Size < cur.Size {
				cur = x
			}
		}
	}
	return cur
}

// TreeSize returns the sum of the sizes of all directories in the tree with a size at most limit.
func TreeSize(node *FSNode, limit int) (total int) {
	if !node.IsDir {
		return
	}

	if node.Size <= limit {
		total += node.Size
	}
	for _, node := range node.Entries {
		total += TreeSize(node, limit)
	}
	return
}
