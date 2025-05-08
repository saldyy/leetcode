package libs

import (
	"fmt"
)

type Item struct {
	value    string
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Print() {
	fmt.Printf("---------------------\n")
	for _, item := range pq {
		if item == nil {
			continue
		}
		fmt.Printf("Value: %s, priority: %d\n", item.value, item.priority)
	}
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)

	item := old[n-1]
	old[n-1] = nil

	*pq = old[0 : n-1]
	return item
}
