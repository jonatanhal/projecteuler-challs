/*
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 * 
 * Find the sum of all the primes below two million.
*/

package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(SumOfPrimes(2000000))
}

func SumOfPrimes(max uint64) uint64 {
	// adapted from p7
	m     := uint64(math.Sqrt(float64(max)))
	n     := uint64(float64(max)/math.Log(float64(max))) + m<<11
	flags := make([]bool,n)
	sum   := uint64(0)

	fmt.Println("m:",m,"n:",n)

	for i := range flags {
		flags[i]=true
	}

	for i := uint64(2); i <= m; i++ {
		if flags[i] {
			x := uint64(0)
			for j := i+i; j < max; j = i*2+i*x {
				x++
				flags[j] = false
			}
		}
	}

	for i := range flags {
		if flags[i] {
			sum += uint64(i)
		}
	}
	return sum
}
