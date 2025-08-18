package chapter8

import "slices"

// SetCovering finds the smallest set of stations that cover all the required states.
// stations: a map where keys are station names and values are sets of states each station covers
// statesNeeded: a set of states that need to be covered
// Returns a slice of station names that together cover all the required states
func SetCovering(stations map[string][]string, statesNeeded []string) []string {
	statesStillNeeded := slices.Clone(statesNeeded)
	stationsUsed := []string{}

	for len(statesStillNeeded) > 0 {
		bestStation := ""
		maxStation := 0
		for station, states := range stations {
			// grab states that are covered by this station from statesStillNeeded
			statesCovered := grabIntersection(statesStillNeeded, states)

			// if statesCovered is bigger that maxStation
			// set bestStation
			if len(statesCovered) > maxStation {
				maxStation = len(statesCovered)
				bestStation = station
			}
		}

		// Check if a valid station was found (maxStation > 0)
		if maxStation == 0 {
			// No station covers any remaining states, can't complete the coverage
			break
		}

		// now that you have best station, rest all those states from statesTillNeeded
		for _, v := range stations[bestStation] {
			index := slices.Index(statesStillNeeded, v)
			if index != -1 {
				statesStillNeeded = slices.Delete(statesStillNeeded, index, index+1)
			}
		}

		// add the station to stationsUsed
		stationsUsed = append(stationsUsed, bestStation)
	}

	return stationsUsed
}

func grabIntersection(array1, array2 []string) []string {
	// Use a map for O(1) lookups
	lookup := make(map[string]struct{})
	for _, item := range array1 {
		lookup[item] = struct{}{}
	}

	// Check which elements from array2 exist in the map
	var result []string
	for _, item := range array2 {
		if _, exists := lookup[item]; exists {
			// If we find an item in both arrays, add it to results
			result = append(result, item)
			// Optional: delete from map to handle duplicates
			// delete(lookup, item)
		}
	}

	return result
}
