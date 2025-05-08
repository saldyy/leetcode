package problem3341

import (
	"container/heap"
	"fmt"
	"math"
)

var (
	dx = []int{0, 1, -1, 0}
	dy = []int{1, 0, 0, -1}
)

type Point struct {
	x, y, t int
}

type Item struct {
	value    Point
	priority int
}

type PriorityQueue []*Item

func (p Point) isValid(n, m int) bool {
	return p.x >= 0 && p.y >= 0 && p.x < n && p.y < m
}

func (pq PriorityQueue) Print() {
	fmt.Printf("---------------------\n")
	for _, item := range pq {
		if item == nil {
			continue
		}
		fmt.Printf("Value: %v, priority: %d\n", item.value, item.priority)
	}
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printArr(arr [][]int) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("%3d ", arr[i][j])
		}
		fmt.Println()
	}
}

func minTimeToReach(moveTime [][]int) int {
	n := len(moveTime)
	m := len(moveTime[0])
	visited := make([][]int, len(moveTime))
	for i := range moveTime {
		visited[i] = make([]int, len(moveTime[i]))
		for j := range moveTime[i] {
			visited[i][j] = math.MaxInt
		}
	}

	var pq PriorityQueue
	heap.Init(&pq)
	heap.Push(&pq, &Item{value: Point{x: 0, y: 0, t: 0}, priority: 0})
	visited[0][0] = 0

	for pq.Len() != 0 {
		top := heap.Pop(&pq).(*Item)
		curTime := top.priority
		curPoint := top.value

		for i := 0; i < 4; i++ {
			nextPoint := Point{
				x: curPoint.x + dx[i],
				y: curPoint.y + dy[i],
				t: curPoint.t + 1,
			}
			if !nextPoint.isValid(n, m) {
				continue
			}

			if visited[nextPoint.x][nextPoint.y] != math.MaxInt {
				continue
			}

			inc := 1
			if nextPoint.t%2 == 0 {
				inc = 2
			}

			nextTime := max(moveTime[nextPoint.x][nextPoint.y], curTime) + inc

			if nextTime < visited[nextPoint.x][nextPoint.y] {
				visited[nextPoint.x][nextPoint.y] = nextTime
				heap.Push(&pq, &Item{value: nextPoint, priority: nextTime})
			}
		}

	}

	return visited[n-1][m-1]
}

func Execute(a [][]int) int {

	return minTimeToReach(a)
}
