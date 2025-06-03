package chapter4

func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	} else {
		return numbers[0] + Sum(numbers[1:])
	}
}

func Count(items []int) int {
	if len(items) == 0 {
		return 0
	} else {
		return 1 + Count(items[1:])
	}
}

func Max(items []int) int {
	if len(items) == 0 {
		return 0
	} else if len(items) == 1 {
		return items[0]
	} else {
		maxRest := Max(items[1:])

		if items[0] > maxRest {
			return items[0]
		} else {
			return maxRest
		}
	}
}

func BinarySearch(array []int, n int) int {
	if len(array) == 0 {
		return -1
	} else {

		mid := len(array) / 2
		guess := array[mid]

		if guess == n {
			return mid
		} else if guess > n {
			return BinarySearch(array[:mid], n)
		} else {
			res := BinarySearch(array[mid+1:], n)

			if res == -1 {
				return -1
			} else {
				return mid + 1 + res
			}
		}
	}
}
