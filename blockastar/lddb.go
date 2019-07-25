package blockastar

import (
	"container/heap"
	"fmt"
)

// LDDB Type
type LDDB map[*Cell]map[*Cell]float64

// Print LDDB
func (lddb LDDB) Print() {
	for k1, v1 := range lddb {
		for k2, v2 := range v1 {
			fmt.Println(k1, k2, v2)
		}
	}
}

// GenerateLDDB from a Grid
func GenerateLDDB(grid Grid) LDDB {
	lddb := make(LDDB)
	for i := 0; i < grid.R/grid.K; i++ {
		for j := 0; j < grid.C/grid.K; j++ {
			block := grid.Blocks[i][j]
			for ii := 0; ii < grid.K; ii++ {
				for jj := 0; jj < grid.K; jj++ {
					if (ii == 0) || (jj == 0) || (ii == grid.K-1) || (jj == grid.K-1) {
						start := block.cells[ii][jj]
						if start.isEmpty {
							UpdateLDDB(lddb, block, start)
						}
					}
				}
			}
		}
	}
	return lddb
}

// UpdateLDDB update the lddb by running Dijkstra algorithm on block from start cell
func UpdateLDDB(lddb LDDB, block *Block, start *Cell) {
	dist := make(map[*Cell]float64)
	for i := 0; i < block.K; i++ {
		for j := 0; j < block.K; j++ {
			cell := block.cells[i][j]
			cell.inQueue = false
			dist[cell] = INF
		}
	}

	start.inQueue = true
	start.heapvalue = 0.0
	dist[start] = 0.0
	pq := make(CellPQ, 1)
	pq[0] = start
	heap.Init(&pq)
	for pq.Len() > 0 {
		currCell := heap.Pop(&pq).(*Cell)
		for _, nextCell := range currCell.sides {
			if dist[currCell]+currCell.Dist(nextCell) < dist[nextCell] {
				dist[nextCell] = dist[currCell] + currCell.Dist(nextCell)
				if !nextCell.inQueue {
					nextCell.heapvalue = dist[nextCell]
					nextCell.inQueue = true
					heap.Push(&pq, nextCell)
				} else {
					pq.Update(nextCell, dist[nextCell])
				}
			}
		}
	}
	lddb[start] = dist
}
