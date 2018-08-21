/* 
Starting in the top left corner of a 2×2 grid, and only being able to
move to the right and down, there are exactly 6 routes to the bottom
right corner.

How many such routes are there through a 20×20 grid?
---
Note:

After some googling on lattice paths (the name of the challenge), I'm
to understand that the amount of possible paths is represented by a
'binomial' algorithm/equation, for which golang provides
functionality.
  */



package main

import (
	"fmt"
	"math/big"
)

func main() {
	height := int64(20)
	i := big.NewInt(int64(0))
	i.Binomial(height*2,height)
	fmt.Printf("%s\n",i.String())
}
