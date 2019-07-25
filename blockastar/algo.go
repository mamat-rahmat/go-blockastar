package blockastar

import (
	"container/heap"
	"math"
)

func initialize(cell *Cell, lddb LDDB) *Block {
	UpdateLDDB(lddb, cell.block, cell)
	return cell.block
}

func initStart(block *Block, start *Cell, lddb LDDB) {
	block.heapvalue = 0.0
	for i := 0; i < block.K; i++ {
		for j := 0; j < block.K; j++ {
			if (i == 0) || (j == 0) || (i == block.K-1) || (j == block.K-1) {
				cell := block.cells[i][j]
				if cell.isEmpty {
					cell.g = lddb[start][cell]
					block.ingress[cell] = void{}
				}
			}
		}
	}
}

// Run Block A* Algorithm
func Run(start *Cell, goal *Cell, lddb LDDB) float64 {
	startBlock := initialize(start, lddb)
	initStart(startBlock, start, lddb)
	goalBlock := initialize(goal, lddb)
	if startBlock == goalBlock {
		return lddb[start][goal]
	}

	length := INF

	pq := make(BlockPQ, 1)
	pq[0] = startBlock
	heap.Init(&pq)
	// totaltotal := 0
	for pq.Len() > 0 {
		currBlock := heap.Pop(&pq).(*Block)
		currBlock.inQueue = false
		if currBlock.heapvalue >= length {
			break
		}
		Y := currBlock.ingress
		if currBlock == goalBlock {
			for y := range Y {
				length = math.Min(length, y.g+lddb[y][goal])
			}
		}
		total := 0
		for _, side := range currBlock.sides {
			nextBlock := side.neighbour
			newheapvalue := INF
			for _, x := range side.egress {
				xin := x.in
				for y := range Y {
					xin.g = math.Min(xin.g, y.g+lddb[y][xin])
					total++
				}
				for _, xout := range x.out {
					newxoutg := xin.g + xin.Dist(xout)
					if newxoutg < xout.g {
						xout.g = newxoutg
						nextBlock.ingress[xout] = void{}
						newheapvalue = math.Min(newheapvalue, xout.g+xout.Dist(goal))
					}
				}
			}
			if newheapvalue < nextBlock.heapvalue {
				if !nextBlock.inQueue {
					nextBlock.inQueue = true
					nextBlock.heapvalue = newheapvalue
					heap.Push(&pq, nextBlock)
				} else {
					pq.Update(nextBlock, newheapvalue)
				}
			}
		}
		currBlock.ingress = map[*Cell]void{}
	}
	return length
}
