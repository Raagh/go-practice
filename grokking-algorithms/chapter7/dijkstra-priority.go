package chapter7

import (
	"container/heap"
	"math"
)

// I didnt created this, this was copilot, lets compare and learn priority queue
func DijkstraPriority(graph map[string]map[string]int, start string, end string) ([]string, int) {
	if _, ok := graph[start]; !ok {
		return []string{}, -1
	}
	if _, ok := graph[end]; !ok {
		return []string{}, -1
	}

	costs := map[string]int{}
	parents := map[string]string{}
	visited := map[string]bool{}

	for node := range graph {
		costs[node] = math.MaxInt32
	}
	costs[start] = 0

	// Priority queue: slice of [cost, node]
	pq := &minHeap{}
	heap.Init(pq)
	heap.Push(pq, &item{node: start, cost: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*item)
		node := curr.node
		if visited[node] {
			continue
		}
		visited[node] = true

		for neighbor, neighborCost := range graph[node] {
			newCost := costs[node] + neighborCost
			if newCost < costs[neighbor] {
				costs[neighbor] = newCost
				parents[neighbor] = node
				heap.Push(pq, &item{node: neighbor, cost: newCost})
			}
		}
	}

	if costs[end] == math.MaxInt32 {
		return []string{}, -1
	}

	// reconstruct path
	path := []string{}
	curr := end
	for curr != "" {
		path = append(path, curr)
		curr = parents[curr]
	}
	reverse(path)

	return path, costs[end]
}

type item struct {
	node string
	cost int
}

type minHeap []*item

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(*item))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
