package chapter6

import (
	"reflect"
	"slices"
	"testing"
)

func TestBreadFirstSearch(t *testing.T) {
	tests := []struct {
		name     string
		graph    map[string][]string
		start    string
		expected []string // allow multiple valid answers
	}{
		{
			name: "Finds seller(thom)",
			graph: map[string][]string{
				"you":    {"alice", "bob", "claire"},
				"bob":    {"anuj", "peggy"},
				"alice":  {"you", "peggy", "alice"},
				"claire": {"jonny", "thom", "bob"},
				"anuj":   {},
				"peggy":  {"you"},
				"thom":   {},
				"jonny":  {"peggy", "peggy"},
				"jim":    {"thom"},
			},
			start:    "you",
			expected: []string{"thom"},
		},
		{
			name: "Target not found",
			graph: map[string][]string{
				"a": {"b"},
				"b": {"c"},
				"c": {},
			},
			start:    "a",
			expected: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsThereASeller(tt.graph, tt.start)
			found := slices.Contains(tt.expected, result)
			if !found {
				t.Errorf("Expected one of %v, got '%s'", tt.expected, result)
			}
		})
	}
}

func TestShortestPathToSeller(t *testing.T) {
	// Example graph
	graph := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}

	expected := []string{"you", "claire", "thom"}
	result := ShortestPathToSeller(graph, "you")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
