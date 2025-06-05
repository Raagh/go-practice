package chapter4

func Sum(array []int) int {
	if len(array) == 0 {
		return 0
	} else {
		return array[0] + Sum(array[1:])
	}
}

func Count(array []int) int {
	if len(array) == 0 {
		return 0
	} else {
		return 1 + Count(array[1:])
	}
}

func Max(array []int) int {
	if len(array) == 0 {
		return 0
	} else if len(array) == 1 {
		return array[0]
	} else {
		maxRes := Max(array[1:])
		if array[0] > maxRes {
			return array[0]
		} else {
			return maxRes
		}
	}
}

func BinarySearch(array []int, n int) int {
	if len(array) == 0 {
		return -1
	}

	mid := len(array) / 2
	guess := array[mid]

	if n == guess {
		return mid
	} else if n < guess {
		return BinarySearch(array[:mid], n)
	} else {
		res := BinarySearch(array[mid+1:], n)

		if res == -1 {
			return -1
		}

		return mid + 1 + res
	}
}
