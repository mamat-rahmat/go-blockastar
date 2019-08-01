package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mamat-rahmat/go-blockastar/blockastar"
)

func main() {
	args := os.Args[1:]
	maps := args[0]
	K, _ := strconv.Atoi(args[1])
	x1, _ := strconv.Atoi(args[2])
	y1, _ := strconv.Atoi(args[3])
	x2, _ := strconv.Atoi(args[4])
	y2, _ := strconv.Atoi(args[5])
	grid := blockastar.BuildGridFromMaps("maps/"+maps+".map", K)
	lddb := blockastar.GenerateLDDB(grid)
	start := time.Now()
	len := blockastar.Run(&grid, grid.Cells[x1][y1], grid.Cells[x2][y2], lddb)
	duration := time.Since(start)
	fmt.Println("length : ", len)
	fmt.Println("duration : ", int64(duration))
}
