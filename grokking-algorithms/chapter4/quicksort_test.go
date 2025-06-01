package chapter4

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse order", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"random order", []int{3, 1, 4, 5, 2}, []int{1, 2, 3, 4, 5}},
		{"empty slice", []int{}, []int{}},
		{"single element", []int{42}, []int{42}},
		{"duplicates", []int{2, 3, 2, 1, 3}, []int{1, 2, 2, 3, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, len(tt.input))
			copy(got, tt.input)
			got = QuickSort(got)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("QuickSort(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}
