package blockastar

import (
	"fmt"
)

// Egress type
type Egress struct {
	in  *Cell
	out []*Cell
}

type void struct{}

// Block type
type Block struct {
	index     int
	x         int
	y         int
	K         int
	inQueue   bool
	heapvalue float64
	cells     [][]*Cell
	ingress   map[*Cell]void
	sides     []struct {
		neighbour *Block
		egress    []Egress
	}
}

// String representation of Block
func (b *Block) String() string {
	return fmt.Sprintf("<Block (%d,%d) %.3f>", b.x+1, b.y+1, b.heapvalue)
}
