package chapter7

func Dijkstra(graph map[string]map[string]int, start string, end string) (path []string, weight int) {
	costs := map[string]int{}
	parents := map[string]string{}
	visited := map[string]bool{}

	// populate costs and parents
	for name, cost := range graph[start] {
		costs[name] = cost
		parents[name] = start
	}
	costs[end] = int(^uint(0) >> 1)

	node := findLowestCostNode(costs, visited)

	for node != "" {
		neighbors := graph[node]
		for neighbor, neighborCost := range neighbors {
			newCost := costs[node] + neighborCost
			if c, ok := costs[neighbor]; !ok || newCost < c {
				costs[neighbor] = newCost
				parents[neighbor] = node
			}
		}
		visited[node] = true
		node = findLowestCostNode(costs, visited)
	}

	// reconstruct path
	path = []string{}
	curr := end
	for {
		path = append(path, curr)
		p, ok := parents[curr]
		if !ok {
			break
		}

		curr = p
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
	lowestNodeCost := int(^uint(0) >> 1)
	lowestNodeCostName := ""
	for name, cost := range costs {
		if visited[name] {
			continue
		}

		if cost < lowestNodeCost {
			lowestNodeCost = cost
			lowestNodeCostName = name
		}
	}

	return lowestNodeCostName
}
