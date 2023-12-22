package main

import "fmt"

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime number, by definition", n)
	}

	if n < 0 {
		return false, "negative number is not prime number, by definition"
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime number because it is devisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is prime number", n)
}
