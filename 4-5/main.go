package main

func Fibo(N int) int {
	if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	return Fibo(N-1) + Fibo(N-2)
}
