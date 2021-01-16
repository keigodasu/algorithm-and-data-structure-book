package main

func MultipleArray(A, B []int) int {
	sum := 0

	for i := len(A) - 1; i >= 0; i-- {
		A[i] += sum
		amari := A[i] % B[i]
		D := 0

		if amari != 0 {
			D = B[i] - amari
		}

		sum += D
	}

	return sum
}
