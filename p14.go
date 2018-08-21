
package main

import (
	"fmt"
)

/*
if the previous term is even, the next term is one half the previous
term. If the previous term is odd, the next term is 3 times the
previous term plus 1.
*/
func collatz(n uint64) uint64 {
	if n % 2 == 0 {
		return n / 2
	} else {
		// bitshifting is faster(?) than doing straight
		// multiples
		return (n<<1)+n+1
	}
}

func main() {
	n := uint64(0)
	h := 0
	c := []uint64{}
	for i := uint64(13); i < 1000000; i++ {
		s := []uint64{}
		nn := i
		for ; nn != 1 ; {
			nn = collatz(nn)
			s = append(s, nn)
		}
		if len(s) > h {
			//fmt.Printf("New highscore!\nchain{%#v}\nstart{%d}\nlen(chain)->%d\n",s,i,len(s))
			h = len(s)
			n = i
			c = append([]uint64{},s...)
		}
		//fmt.Printf("Debug sequence: %#v\n",s)
	}
	fmt.Printf("winner: %d\nSequence: %v\n",n,c)
}
