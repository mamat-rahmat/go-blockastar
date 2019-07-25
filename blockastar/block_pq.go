package blockastar

import "container/heap"

// A BlockPQ implements heap.Interface and holds Block.
type BlockPQ []*Block

// Len operator
func (pq BlockPQ) Len() int { return len(pq) }

// Less operator
func (pq BlockPQ) Less(i, j int) bool {
	return pq[i].heapvalue < pq[j].heapvalue
}

// Swap operation
func (pq BlockPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push operation
func (pq *BlockPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Block)
	item.index = n
	item.inQueue = true
	*pq = append(*pq, item)
}

// Pop operation
func (pq *BlockPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	item.inQueue = false
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an Block in the queue.
func (pq *BlockPQ) Update(item *Block, heapvalue float64) {
	item.heapvalue = heapvalue
	heap.Fix(pq, item.index)
}
