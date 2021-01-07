package main

func main() {

}

func FindSum(list []int, num int) bool {
	exist := false

	for bit := 0; bit < (1 << len(list)); bit++ {
		sum := 0
		for i := 0; i < len(list); i++ {
			if (bit>>i)&1 == 1 {
				sum += list[i]
			}
		}

		if sum == num {
			exist = true
		}
	}

	return exist
}
