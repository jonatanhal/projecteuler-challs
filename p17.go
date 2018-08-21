/*
If the numbers 1 to 5 are written out in words: one, two, three, four,
five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were
written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred
and forty-two) contains 23 letters and 115 (one hundred and fifteen)
contains 20 letters. The use of "and" when writing out numbers is in
compliance with British usage.  */

package main

import (
	"strings"
	"fmt"
)

var foo = map[int]string {
	1:"one",
	2:"two",
	3:"three",
	4:"four",
	5:"five",
	6:"six",
	7:"seven",
	8:"eight",
	9:"nine",
	10:"ten",
	11:"eleven",
	12:"twelve",
	13:"thirteen",
	14:"fourteen",
	15:"fifteen",
	16:"sixteen",
	17:"seventeen",
	18:"eighteen",
	19:"nineteen",
	20:"twenty",
	30:"thirty",
	40:"forty",
	50:"fifty",
	60:"sixty",
	70:"seventy",
	80:"eighty",
	90:"ninety",
	100:"one hundred",
	1000:"one thousand",
}

func translate(n int) string {
	const hndrd = "hundred"

	w, ok := foo[n]
	if ok {
		return w
	} else {
		switch {
		case n < 1000 && n > 100:
			b := n / 100
			if n % 100 != 0 {
				t := translate(n % 100)
				return fmt.Sprintf("%s %s and %s",foo[b],hndrd,t)
			} else {
				return fmt.Sprintf("%s %s",foo[b],hndrd)
			}
			
		case n < 100:
			r := n % 10
			ww, ok := foo[r]
			if ! ok {
				panic("Segmentation fault")
			}
			www, ok := foo[n-r]
			if ! ok {
				panic("Fatal: I/O Error")
			}
			return fmt.Sprintf("%s-%s",www,ww)
		default:
			panic("Permission denied")
		}
	}
	return ""
}

func slen(s string) int {
	// challange calls for not counting whitespace or "and"
	ss := strings.Replace(s," and ","",-1)
	ss = strings.Replace(ss," ","",-1)
	ss = strings.Replace(ss,"-","",-1)
	return len(ss)
}

func main() {
	sum:=0
	for i := 1; i < 1001; i++ {
		s := translate(i)
		sum += slen(s)
		fmt.Printf("%s, ",s)
	}
	fmt.Printf("\nTotal string length: %d\n",sum)
}
