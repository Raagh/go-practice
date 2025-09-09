package chapter8

import (
	"reflect"
	"sort"
	"testing"
)

func TestMaxClassesScheduled(t *testing.T) {
	tests := []struct {
		name     string
		classes  []Class
		expected int // expected number of classes
	}{
		{
			name: "Basic example from book",
			classes: []Class{
				{Start: 9, End: 10},  // Art
				{Start: 9, End: 11},  // English
				{Start: 10, End: 11}, // Math
				{Start: 11, End: 12}, // CS
				{Start: 11, End: 13}, // Music
				{Start: 13, End: 14}, // PE
			},
			expected: 4, // Optimal solution: Art, Math, CS, PE
		},
		{
			name: "No classes",
			classes: []Class{},
			expected: 0,
		},
		{
			name: "One class",
			classes: []Class{
				{Start: 1, End: 2},
			},
			expected: 1,
		},
		{
			name: "All classes overlap",
			classes: []Class{
				{Start: 1, End: 5},
				{Start: 2, End: 6},
				{Start: 3, End: 7},
			},
			expected: 1,
		},
		{
			name: "No overlaps",
			classes: []Class{
				{Start: 1, End: 2},
				{Start: 3, End: 4},
				{Start: 5, End: 6},
			},
			expected: 3,
		},
		{
			name: "Complex overlapping classes",
			classes: []Class{
				{Start: 1, End: 3},
				{Start: 2, End: 4},
				{Start: 3, End: 5},
				{Start: 4, End: 6},
				{Start: 5, End: 7},
				{Start: 6, End: 8},
				{Start: 7, End: 9},
				{Start: 8, End: 10},
			},
			expected: 4, // Optimal: {1,3}, {3,5}, {5,7}, {7,9}
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MaxClassesScheduled(test.classes)

			// Check if the number of classes is as expected
			if len(result) != test.expected {
				t.Errorf("Expected %d classes, got %d", test.expected, len(result))
			}

			// Verify there are no overlapping classes in the result
			for i := 0; i < len(result); i++ {
				for j := i + 1; j < len(result); j++ {
					if (result[i].Start < result[j].End && result[i].End > result[j].Start) ||
						(result[j].Start < result[i].End && result[j].End > result[i].Start) {
						t.Errorf("Overlapping classes found: %+v and %+v", result[i], result[j])
					}
				}
			}

			// For the basic example, check if it's one of the optimal solutions
			if test.name == "Basic example from book" {
				// One possible optimal solution
				possibleOptimal := []Class{
					{Start: 9, End: 10},  // Art
					{Start: 10, End: 11}, // Math
					{Start: 11, End: 12}, // CS
					{Start: 13, End: 14}, // PE
				}

				// Sort both slices to make comparison possible
				sortClasses := func(classes []Class) {
					sort.Slice(classes, func(i, j int) bool {
						if classes[i].Start != classes[j].Start {
							return classes[i].Start < classes[j].Start
						}
						return classes[i].End < classes[j].End
					})
				}

				sortClasses(result)
				sortClasses(possibleOptimal)

				if reflect.DeepEqual(result, possibleOptimal) {
					t.Logf("Found the expected class combination")
				} else {
					t.Logf("Found a different class combination: %+v", result)
					// Note: This is not an error since there could be multiple valid solutions
				}
			}
		})
	}
}

// This is a helper function to verify if a solution is valid (no overlapping classes)
// It's not used in the tests above but could be useful for debugging
func validateNoOverlaps(classes []Class) bool {
	for i := 0; i < len(classes); i++ {
		for j := i + 1; j < len(classes); j++ {
			// Check if classes overlap
			if (classes[i].Start < classes[j].End && classes[i].End > classes[j].Start) ||
				(classes[j].Start < classes[i].End && classes[j].End > classes[i].Start) {
				return false
			}
		}
	}
	return true
}

