package util

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Factorial(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}
