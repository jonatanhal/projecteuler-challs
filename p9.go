/* 
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a2 + b2 = c2

For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import "fmt"

func main() {
	max := 1000

	for n := 1 ; n < max ; n++ {
		for m := 1 ; m < max ; m++ {
			a := m*m - n*n
			b := m*n<<1
			c := m*m + n*n
			if a < 0 || b < 0 || c < 0 {
				continue
			}
			if a*a + b*b == c*c {
				if a+b+c == 1000 {
					fmt.Printf("Winner:\na: %d, b: %d, c: %d\n",a,b,c)
					break
				}
			}
		}
	}
}
