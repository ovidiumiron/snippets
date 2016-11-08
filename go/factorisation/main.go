package main

import (
	"fmt"
	"os"
)

type Information struct {
	prime int
	count int
}

func next_prime(n int, primes []int) []int {
	for _, prime := range primes {
		if n%prime == 0 {
			return next_prime(n+1, primes)
		}
	}
	return append(primes, n)
}

func main() {
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		os.Exit(1)
	}
	factors := make([]Information, 0)
	primes := []int{2}
	var prime int
	for input > 1 {
		prime = primes[len(primes)-1]
		if input%prime == 0 {
			if len(factors) == 0 {
				factors = append(factors, Information{prime: prime, count: 1})
			} else {
				last := factors[len(factors)-1]
				if last.prime == prime {
					last.count += 1
					factors[len(factors)-1] = last
				} else {
					factors = append(factors, Information{prime: prime, count: 1})
				}
			}
			input = input / prime
		} else {
			primes = next_prime(prime+1, primes)
		}
	}

	fmt.Println(factors)
}
