package chapter3

type Item any

type Key struct{}

type Box struct {
	Items []Item
}

func FindKey(box *Box) bool {
	boxes := []Item{box}

	for _, item := range boxes {
		switch v := item.(type) {
		case *Key: // Check if item is a pointer to Key
			return true
		case *Box: // Check if item is a pointer to Box
			boxes = append(boxes, v.Items...) // Add the items inside the box to the search
		}
	}
	return false
}
