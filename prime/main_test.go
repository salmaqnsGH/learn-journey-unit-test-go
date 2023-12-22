package main

import (
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"PRIME", 7, true, "7 is prime number"},
		{"NOT PRIME", 21, false, "21 is not prime number because it is devisible by 3"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected '%s' but got '%s'", e.name, e.msg, msg)
		}
	}
}
