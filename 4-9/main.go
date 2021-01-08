package main

func PartialSum(i, w int, list []int) bool {
	if i == 0 {
		return w == 0
	}

	if PartialSum(i-1, w, list) {
		return true
	}

	if PartialSum(i-1, w-list[i-1], list) {
		return true
	}

	return false
}
