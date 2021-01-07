package main

func main() {

}

func FindMinimumPair(list []int, K int) int {
	min := 2000000
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1; j++ {
			if list[i]+list[j] < K {
				continue
			}
			if list[i]+list[j] < min {
				min = list[i] + list[j]
			}
		}
	}

	return min
}
