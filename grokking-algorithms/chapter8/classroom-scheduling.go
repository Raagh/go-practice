package chapter8

import "sort"

// Class represents a class with start and end time
type Class struct {
	Start int // Start time
	End   int // End time
}

// MaxClassesScheduled finds the maximum number of classes that can be scheduled in a single classroom
// without any overlap, and returns the list of selected classes
// Input: a slice of Class structs, each with Start and End times
// Output: a slice containing the selected classes
func MaxClassesScheduled(classes []Class) []Class {
	if len(classes) == 0 {
		return nil
	}

	sort.Slice(classes, func(i, j int) bool {
		return classes[i].End < classes[j].End
	})

	firstClass := classes[0]
	selectedClasses := []Class{firstClass}
	next := findNext(classes, firstClass.End)
	if next.Start != 0 {
		selectedClasses = append(selectedClasses, next)
	}
	for next.Start > 0 {
		lastClass := selectedClasses[len(selectedClasses)-1]
		next = findNext(classes, lastClass.End)
		if next.Start != 0 {
			selectedClasses = append(selectedClasses, next)
		}
	}

	return selectedClasses
}

func findNext(classes []Class, end int) Class {
	next := Class{}
	for _, class := range classes {
		if class.Start >= end {
			next = class
			break
		}
	}

	return next
}
