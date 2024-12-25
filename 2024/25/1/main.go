package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	schematics := strings.Split(io.ReadAll(), "\n\n")
	var locks [][5]int
	var keys [][5]int
	for _, schematic := range schematics {
		rows := strings.Split(strings.TrimSpace(schematic), "\n")
		if rows[0] == "#####" {
			lock := [5]int{}
			for i := 1; i < len(rows); i++ {
				for j := 0; j < 5; j++ {
					if rows[i][j] == '#' {
						lock[j]++
					}
				}
			}
			locks = append(locks, lock)
		} else if rows[len(rows)-1] == "#####" {
			key := [5]int{}
			for i := len(rows) - 2; i > 0; i-- {
				for j := 0; j < 5; j++ {
					if rows[i][j] == '#' {
						key[j]++
					}
				}
			}
			keys = append(keys, key)
		}
	}
	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			good := true
			for i := 0; i < 5; i++ {
				if lock[i]+key[i] > 5 {
					good = false
					break
				}
			}
			if good {
				count++
			}
		}
	}
	fmt.Println("count:", count)
}
