package main

//this is reference code
func P(x int) bool {
	return true
}

func binarySerach(left, right int) int {
	var mid int
	for right-left > 1 {
		mid = left + ((right - left) / 2)
		if P(mid) {
			right = mid
		} else {
			left = mid
		}
	}

	return right
}
