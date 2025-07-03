package chapter6

import "slices"

func isSeller(current string) bool {
	return current[len(current)-1] == 'm'
}

func IsThereASeller(graph map[string][]string, start string) string {
	queue := []string{start}
	visited := map[string]bool{}

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]

		if visited[element] {
			continue
		}

		if isSeller(element) {
			return element
		}

		queue = append(queue, graph[element]...)
		visited[element] = true
	}

	return ""
}

func ShortestPathToSeller(graph map[string][]string, start string) []string {
	type Node struct {
		Name string
		Path []string
	}
	queue := []Node{{Name: start, Path: []string{start}}}
	visited := map[string]bool{}

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]

		if visited[element.Name] {
			continue
		}

		if isSeller(element.Name) {
			return element.Path
		}

		for _, neighbor := range graph[element.Name] {
			if visited[neighbor] {
				continue
			}

			newPath := slices.Clone(element.Path)
			newPath = append(newPath, neighbor)
			queue = append(queue, Node{Name: neighbor, Path: newPath})
		}

		visited[element.Name] = true
	}

	return nil
}
