/*
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we
 * can see that the 6th prime is 13.
 * 
 * What is the 10 001st prime number?
*/

package main

import (
	"fmt"
	"math"
)

/*
 Input: an integer n > 1.
 
 Let A be an array of Boolean values, indexed by integers 2 to n,
 initially all set to true.
 
 for i = 2, 3, 4, ..., not exceeding √n:
   if A[i] is true:
     for j = i2, i2+i, i2+2i, i2+3i, ..., not exceeding n:
       A[j] := false.
 
 Output: all i such that A[i] is true.
*/
func PrimeAtIndex(n uint64) uint64 {
	flags  := make([]bool,n<<8)
	max    := uint64(math.MaxUint64)
	m      := uint64(math.Sqrt(float64(max)))
	count  := 0

	outer:
	for i := uint64(2); i <= m; i++ {
		if i >= n {
			break
		}
		if !flags[i] {
			x := uint64(0)
			for j := i+i; j < max; j = i*2+i*x {
				x++
				if j >= n {
					continue outer
				} else {
					flags[j] = true
				}
			}
		}
	}

	for i := range flags {
		if !flags[i] {
			count++
			if uint64(count) == n {
				return uint64(i)
			}
		}
	}
	return 0
}

func main() {
	n := uint64(10001)
	primes := PrimeAtIndex(n)
	fmt.Println(primes)
}

