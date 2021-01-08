package main

func Gcd(m, n int) int {
	//base caee
	if n == 0 {
		return m
	}

	return Gcd(n, m%n)
}
