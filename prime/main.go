package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	intro()

	// create a channel to indicate when the user wants to quit
	doneChan := make(chan bool)

	// start a goroutine to read user input
	go readUserInput(doneChan)

	// block until the doneChan gets value
	<-doneChan

	// close the channel
	close(doneChan)

	// say goodbye
	fmt.Println("Goodbye.")
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check to see if user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convers what the user typed into an int
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter the whole number", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

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

func intro() {
	fmt.Println("Is it prime?")
	fmt.Println("------------")
	fmt.Println("Enter a number, and we will tell you if it is prime number")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}
