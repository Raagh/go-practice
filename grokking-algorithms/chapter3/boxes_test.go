package chapter3

import "testing"

func TestFindKey(t *testing.T) {
	// Test case 1: Key is present in a nested box
	key := &Key{}
	box3 := &Box{Items: []Item{key}}
	box2 := &Box{Items: []Item{}}
	box1 := &Box{Items: []Item{box2, box3}}

	if !FindKey(box1) {
		t.Errorf("Expected to find the key, but it was not found")
	}

	// // Test case 2: Key is not present
	// box4 := &Box{Items: []Item{}}
	// box5 := &Box{Items: []Item{box4}}
	//
	// if FindKey(box5) {
	// 	t.Errorf("Expected not to find the key, but it was found")
	// }
	//
	// // Test case 3: Key is in the top-level box
	// box6 := &Box{Items: []Item{key}}
	//
	// if !FindKey(box6) {
	// 	t.Errorf("Expected to find the key in the top-level box, but it was not found")
	// }
	//
	// // Test case 4: Empty box
	// box7 := &Box{Items: []Item{}}
	//
	// if FindKey(box7) {
	// 	t.Errorf("Expected not to find the key in an empty box, but it was found")
	// }
}
