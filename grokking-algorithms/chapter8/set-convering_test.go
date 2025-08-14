package chapter8

import (
	"reflect"
	"sort"
	"testing"
)

func TestSetCovering(t *testing.T) {
	// We want to broadcast to these states
	statesNeeded := []string{
		"mt",
		"wa",
		"or",
		"id",
		"nv",
		"ut",
		"ca",
		"az",
	}

	// List of stations and the states they cover
	stations := map[string][]string{
		"kone": {
			"id",
			"nv",
			"ut",
		},
		"ktwo": {
			"wa",
			"id",
			"mt",
		},
		"kthree": {
			"or",
			"nv",
			"ca",
		},
		"kfour": {
			"nv",
			"ut",
		},
		"kfive": {
			"ca",
			"az",
		},
	}

	// Call the function
	result := SetCovering(stations, statesNeeded)

	// Since there can be multiple valid solutions, we need to check that:
	// 1. All states are covered
	// 2. The number of stations is minimal (should be 4 in this case)

	// Check if all states are covered
	coveredStates := make(map[string]bool)
	for _, station := range result {
		for _, state := range stations[station] {
			coveredStates[state] = true
		}
	}

	// Verify all required states are covered
	for _, state := range statesNeeded {
		if !coveredStates[state] {
			t.Errorf("State %s not covered by selected stations", state)
		}
	}

	// The optimal solution should use 4 stations
	if len(result) != 4 {
		t.Errorf("Expected 4 stations, got %d", len(result))
	}

	// For a more specific test, if we know one particular optimal solution
	// Note: There might be multiple valid solutions with 4 stations
	expectedStations := []string{"kone", "ktwo", "kthree", "kfive"}
	sort.Strings(result)
	sort.Strings(expectedStations)

	// This is optional, as there could be other valid combinations
	if reflect.DeepEqual(result, expectedStations) {
		t.Logf("Found the expected station combination: %v", result)
	} else {
		t.Logf("Found a different station combination: %v", result)
		// Check if it's a valid alternative solution
		// Additional validation can be added here
	}
}
