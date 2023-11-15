package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for n := 2; n < max + 1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		
		if n % 2 == 0 {
			continue
		}

		isPrime := true
		for m := 3; float64(m) <= math.Sqrt(float64(n)) + 1; m++ {
			if n % m == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Println(n)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}