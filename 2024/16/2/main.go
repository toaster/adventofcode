package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/2024/16/reindeerolympics"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	maze := reindeerolympics.ParseMaze(io.ReadLines(), 1, 1000)
	fmt.Println(maze.CountBestPathTiles())
}
