package elfdevice

import "fmt"

// FSNode represents a node in the filesystem of the elf device.
type FSNode struct {
	Entries map[string]*FSNode
	IsDir   bool
	Name    string
	Parent  *FSNode
	Size    int
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
