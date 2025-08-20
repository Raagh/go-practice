package chapter8

// SetCovering picks a minimal-ish set of subsets whose union covers `universe`.
// subsets: name -> slice of elements it covers
// Returns the chosen subset names in greedy order.
func SetCovering(stations map[string][]string, statesNeeded []string) []string {
	needed := toSet(statesNeeded)
	chosen := []string{}

	// Preconvert subset slices to sets for speed
	subsetSets := make(map[string]map[string]struct{}, len(stations))
	for name, elems := range stations {
		subsetSets[name] = toSet(elems)
	}

	for len(needed) > 0 {
		var bestStation string
		var bestCoverCount int
		var bestCover map[string]struct{}

		for name, cover := range subsetSets {
			inter := intersect(needed, cover)
			if len(inter) > bestCoverCount {
				bestCoverCount = len(inter)
				bestCover = inter
				bestStation = name
			}
		}

		if bestCoverCount == 0 {
			// No remaining subset covers any uncovered element -> universe not fully coverable
			break
		}

		// Choose the best station and remove its covered elements from needed
		chosen = append(chosen, bestStation)
		for e := range bestCover {
			delete(needed, e)
		}
		// Optional: remove chosen subset to avoid re-choosing
		delete(subsetSets, bestStation)
	}

	// If needed not empty here, the input subsets don't cover the whole universe.
	return chosen
}

func toSet(items []string) map[string]struct{} {
	s := make(map[string]struct{}, len(items))
	for _, x := range items {
		s[x] = struct{}{}
	}
	return s
}

func intersect(a, b map[string]struct{}) map[string]struct{} {
	out := make(map[string]struct{})
	// iterate over smaller set
	if len(a) > len(b) {
		a, b = b, a
	}
	for x := range a {
		if _, ok := b[x]; ok {
			out[x] = struct{}{}
		}
	}
	return out
}
