package chapter7

import (
	"reflect"
	"testing"
)

func TestDijkstra_SimpleGraph(t *testing.T) {
	graph := map[string]map[string]int{
		"A": {"B": 1, "C": 4},
		"B": {"C": 2, "D": 5},
		"C": {"D": 1},
		"D": {},
	}

	// Suppose the expected shortest path from A to D is: A -> B -> C -> D with weight 4
	path, weight := Dijkstra(graph, "A", "D")
	expectedPath := []string{"A", "B", "C", "D"}
	expectedWeight := 4

	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("expected path %v, got %v", expectedPath, path)
	}
	if weight != expectedWeight {
		t.Errorf("expected weight %d, got %d", expectedWeight, weight)
	}
}

func TestDijkstra_Unreachable(t *testing.T) {
	graph := map[string]map[string]int{
		"A": {"B": 1},
		"B": {},
		"C": {},
	}
	path, weight := Dijkstra(graph, "C", "B")
	if path != nil || weight != 0 {
		t.Errorf("expected nil path and 0 weight for unreachable node, got %v and %d", path, weight)
	}
}
