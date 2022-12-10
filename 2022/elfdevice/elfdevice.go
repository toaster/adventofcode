package elfdevice

import (
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

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
