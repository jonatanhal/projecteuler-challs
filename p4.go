/* 
 * A palindromic number reads the same both ways. The largest palindrome
 * made from the product of two 2-digit numbers is 9009 = 91 × 99.
 * Find the largest palindrome made from the product of two 3-digit
 * numbers.
 */

package main

import (
	"fmt"
	"math"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i,j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func palindrome(s string) bool {
	l := float64(len(s))
	idx := int(math.Floor(l/2.0))
	p1 := s[:idx]
	p2 := s[idx:]
	if p1 == reverse(p2) {
		return true
	}
	return false
}

func main() {
	highscore := 0
	for i := 100; i < 999; i++ {
		for j := 100; j < 999; j++ {
			n := i*j
			if palindrome(fmt.Sprintf("%d",n)) && n > highscore {
				highscore = n
			}
		}
	}
	fmt.Println(highscore)
}

