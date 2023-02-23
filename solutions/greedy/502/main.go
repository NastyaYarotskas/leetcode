package main

import (
	"container/heap"
	"sort"
)

func main() {
	println(findMaximizedCapital(3, 0, []int{1, 2, 5, 7, 6, 8, 3, 1}, []int{0, 1, 2, 6, 4, 3, 1, 1}))
}

type Project struct {
	capital int
	profit  int
}

type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	projects := make([]Project, 0, n)
	for i := 0; i < n; i++ {
		projects = append(projects, Project{
			capital: capital[i],
			profit:  profits[i],
		})
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	j := 0
	for i := 0; i < k; i++ {
		for j < n && projects[j].capital <= w {
			heap.Push(&pq, projects[j].profit)
			j++
		}
		if pq.Len() == 0 {
			break
		}
		w += heap.Pop(&pq).(int)
	}

	return w
}
