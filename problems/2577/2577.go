package problem2577

import (
	"container/heap"
	"fmt"
)

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

type Point struct {
	x    int
	y    int
	time int
}

// PriorityQueue implements heap.Interface and holds Items
type PriorityQueue []Point

// Len returns the length of the priority queue
func (pq PriorityQueue) Len() int { return len(pq) }

// Less determines the priority order (lower number means higher priority)
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].time < pq[j].time
}

// Swap swaps two elements in the priority queue
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds a new item to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	p := x.(Point)
	*pq = append(*pq, p)
}

// Pop removes and returns the highest priority item
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]

	return item
}

// NewPriorityQueue creates and initializes a new priority queue
func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}
	heap.Init(pq)
	return pq
}

// IsEmpty checks if the priority queue is empty
func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}

func isValidPoint(p Point, lenX int, lenY int) bool {
	return p.x >= 0 && p.x < lenX && p.y >= 0 && p.y < lenY
}

func printGrid(grid [][]int) {
	for y := range grid {
		for x := range grid[y] {
			fmt.Printf("%3d ", grid[y][x])
		}
		fmt.Println()
	}
}

func minimumTime(grid [][]int) int {
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}
	lenX := len(grid[0])
	lenY := len(grid)
	visited := make([][]int, len(grid))

	for y := range lenY {
		visited[y] = make([]int, lenX)
		for x := range lenX {
			visited[y][x] = -1
		}
	}

	queue := NewPriorityQueue()
	heap.Push(queue, Point{x: 0, y: 0, time: 0})
	visited[0][0] = 0

	fmt.Println("start")
	for queue.Len() != 0 {

		curPoint := heap.Pop(queue).(Point)
		fmt.Printf("Cur point: %v\n", curPoint)

		if curPoint.x == lenX-1 && curPoint.y == lenY-1 {
			return visited[curPoint.y][curPoint.x]
		}

		curTime := visited[curPoint.y][curPoint.x] + 1

		for i := 0; i < 4; i++ {
			nextX := dx[i] + curPoint.x
			nextY := dy[i] + curPoint.y
			nextPoint := Point{
				x: nextX,
				y: nextY,
        time: curPoint.time + 1,
			}


			fmt.Printf("Next point: %v\n", curPoint)
			if !isValidPoint(nextPoint, lenX, lenY) {
				continue
			}
			if visited[nextPoint.y][nextPoint.x] != -1 {
				continue
			}
			nextTime := curTime + 1
			//
			// if grid[nextPoint.y][nextPoint.x] > curTime {
			// 	nextTime = grid[nextPoint.y][nextPoint.x] + 1
			// }

			visited[nextPoint.y][nextPoint.x] = min(nextTime, visited[nextPoint.y][nextPoint.x])
			nextPoint.time = visited[nextPoint.y][nextPoint.x]

			heap.Push(queue, nextPoint)
		}
	}

	return visited[lenY-1][lenX-1]
}

func Execute(grid [][]int) int {
	return minimumTime(grid)
}
