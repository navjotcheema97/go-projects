package main

import (
	"fmt"
	"math"
)

// IsPrime checks if a number is prime using concurrent operations
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))
	numWorkers := 4 // Number of concurrent workers
	results := make(chan bool, numWorkers)

	// Divide work among goroutines
	for worker := 0; worker < numWorkers; worker++ {
		go func(start int) {
			isPrime := true
			// Each worker checks a range of numbers
			for i := start; i <= sqrt; i += numWorkers * 6 {
				if i >= 5 {
					if n%i == 0 || n%(i+2) == 0 {
						isPrime = false
						break
					}
				}
			}
			results <- isPrime
		}(5 + worker*6)
	}

	// Collect results from all workers
	for i := 0; i < numWorkers; i++ {
		if !<-results {
			return false
		}
	}

	return true
} 


func main() {
	// Test some numbers for primality
	testNumbers := []int{2, 3, 17, 97, 1, 4, 15, 100, 104729, 104730}

	fmt.Println("Testing numbers for primality:")
	for _, num := range testNumbers {
		isPrime := IsPrime(num)
		fmt.Printf("%d is prime: %v\n", num, isPrime)
	}
}