/*


2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	for g := uint64(1); g < math.MaxUint64; g++ {
		status := true
		for d := uint64(2); d <= 20; d++ {
			if !status {
				break
			}
			if g % d != 0 {
				status = false
			}
		}

		if status {
			fmt.Println(g)
			break
		}
	}
}
