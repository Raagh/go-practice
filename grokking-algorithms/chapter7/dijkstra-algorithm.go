package chapter7

import "math"

func Dijkstra(graph map[string]map[string]int, start string, end string) ([]string, int) {
	costs := map[string]int{}
	parents := map[string]string{}
	visited := map[string]bool{}

	// initialize all nodes with max int cost
	for node := range graph {
		costs[node] = math.MaxInt32
	}
	costs[start] = 0

	// populate costs and parents for neighbors of start
	for name, cost := range graph[start] {
		costs[name] = cost
		parents[name] = start
	}

	node := findLowestCostNode(costs, visited)
	for node != end {
		for neighbor, neighborCost := range graph[node] {
			newCost := costs[node] + neighborCost
			if newCost < costs[neighbor] {
				costs[neighbor] = newCost
				parents[neighbor] = node
			}
		}

		visited[node] = true
		node = findLowestCostNode(costs, visited)
	}

	if costs[end] == math.MaxInt32 {
		return nil, 0
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

func reverse(nums []string) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func findLowestCostNode(costs map[string]int, visited map[string]bool) string {
	min := math.MaxInt
	var lowest string
	for node, cost := range costs {
		if visited[node] {
			continue
		}

		if cost < min {
			min = cost
			lowest = node
		}
	}
	return lowest
}
