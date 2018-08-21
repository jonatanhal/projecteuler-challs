/*
2**15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2**1000?

----

Notes:

Bit-shifting two to the left is basically the same as doing (more
expensive(?)) power-operations

2<<1 == 2**2
2<<2 == 2**3
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := []int{}
	s := 0
	i := big.NewInt(int64(2))
	i.Lsh(i,1000)
	for _,c := range i.String() {
		x := rune(c)-0x30 // reduce from ascii
		n = append(n,int(x))
		s += int(x)
	}

	fmt.Printf(
		"2<<1000 == %s\n"+
		"n{%v}\n"+
		"Sum{%d}\n",i.String(),n,s)
}
