package blockastar

import (
	"fmt"
	"math"
)

// Cell type
type Cell struct {
	index     int
	x         int
	y         int
	isEmpty   bool
	g         float64
	h         float64
	inQueue   bool
	heapvalue float64
	block     *Block
	sides     []*Cell
}

// Dist return distance of this Cell to another Cell
func (c *Cell) Dist(d *Cell) float64 {
	dx := math.Abs(float64(c.x - d.x))
	dy := math.Abs(float64(c.y - d.y))
	return math.Sqrt(dx*dx + dy*dy)
}

// ManhattanDistance return manhattan distance of this Cell to another Cell
func (c *Cell) ManhattanDistance(d *Cell) float64 {
	dx := math.Abs(float64(c.x - d.x))
	dy := math.Abs(float64(c.y - d.y))
	return dx + dy
}

func (c *Cell) String() string {
	return fmt.Sprintf("<Cell (%d,%d)>", c.x+1, c.y+1)
}
