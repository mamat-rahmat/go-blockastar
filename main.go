package main

import (
	"fmt"
	"time"

	"github.com/mamat-rahmat/go-blockastar/blockastar"
)

func main() {
	x1 := 11
	y1 := 3
	x2 := 2
	y2 := 3
	grid := blockastar.BuildGridFromMaps("data/paper-tile.map", 5)
	lddb := blockastar.GenerateLDDB(grid)
	start := time.Now()
	len := blockastar.Run(grid.Cells[x1][y1], grid.Cells[x2][y2], lddb)
	duration := time.Since(start)
	fmt.Println("length : ", len)
	fmt.Println("duration : ", duration)
}
