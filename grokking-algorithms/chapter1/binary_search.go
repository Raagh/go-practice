package chapter1

func BinarySearch(array []int, num int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := array[mid]

		if num == guess {
			return mid
		} else if num > guess {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
