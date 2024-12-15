package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/2024/15/deepsea"
	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	mapAndMovements := strings.Split(io.ReadAll(), "\n\n")
	warehouse := deepsea.ParseLanternfishWarehouse(strings.Split(mapAndMovements[0], "\n"))
	warehouse.MoveRobot(strings.Join(strings.Split(mapAndMovements[1], "\n"), ""))
	warehouse.Print()
	fmt.Println(warehouse.SumUpGPSCoordinates())
}
