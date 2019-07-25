package blockastar

import "container/heap"

// A CellPQ implements heap.Interface and holds Cell.
type CellPQ []*Cell

// Len operator
func (pq CellPQ) Len() int { return len(pq) }

// Less operator
func (pq CellPQ) Less(i, j int) bool {
	return pq[i].heapvalue < pq[j].heapvalue
}

// Swap operation
func (pq CellPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push operation
func (pq *CellPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Cell)
	item.index = n
	*pq = append(*pq, item)
}

// Pop operation
func (pq *CellPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an Cell in the queue.
func (pq *CellPQ) Update(item *Cell, heapvalue float64) {
	item.heapvalue = heapvalue
	heap.Fix(pq, item.index)
}
