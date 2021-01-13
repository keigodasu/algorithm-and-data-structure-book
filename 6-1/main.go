package main

func binarySearch(key int, list []int) int {
	left, right := 0, len(list)-1
	var mid int

	for right >= left {
		mid = left + (right-left)/2
		if list[mid] == key {
			return mid
		} else if list[mid] > key {
			right = mid - 1
		} else if list[mid] < key {
			left = mid + 1
		}
	}

	return -1
}
