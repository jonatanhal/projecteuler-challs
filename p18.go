/*
By starting at the top of the triangle below and moving to adjacent
numbers on the row below, the maximum total from top to bottom is 23.

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

NOTE: As there are only 16384 routes, it is possible to solve this
problem by trying every route. However, Problem 67, is the same
challenge with a triangle containing one-hundred rows; it cannot be
solved by brute force, and requires a clever method! ;o)
*/

package main

import (
	"fmt"
)

var q = [][]int{
	[]int{75,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,},
	[]int{95, 64,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,},
	[]int{17, 47, 82,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,},
	[]int{18, 35, 87, 10,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,},
	[]int{20,  4, 82, 47, 65,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,},
	[]int{19,  1, 23, 75,  3, 34,  0,  0,  0,  0,  0,  0,  0,  0,  0,},
	[]int{88,  2, 77, 73,  7, 63, 67,  0,  0,  0,  0,  0,  0,  0,  0,},
	[]int{99, 65,  4, 28,  6, 16, 70, 92,  0,  0,  0,  0,  0,  0,  0,},
	[]int{41, 41, 26, 56, 83, 40, 80, 70, 33,  0,  0,  0,  0,  0,  0,},
	[]int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29,  0,  0,  0,  0,  0,},
	[]int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14,  0,  0,  0,  0,},
	[]int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57,  0,  0,  0,},
	[]int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48,  0,  0,},
	[]int{63, 66,  4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31,  0,},
	[]int{ 4, 62, 98, 27, 23,  9, 70, 98, 73, 93, 38, 53, 60,  4, 23,},
}

func compute() {
	for x := len(q)-2; x >= 0; x-- {
		for y := 0; y <= len(q[0])-2; y++ {
			//fmt.Printf("accessing q[%d][%d]\n",x+1,y+1)
			if q[x+1][y] > q[x+1][y+1] {
				q[x][y] += q[x+1][y];
			} else {
				q[x][y] += q[x+1][y+1];
			}
		}
	}
	fmt.Printf("sum: %d\n",q[0][0])
}

func main() {
	compute()
}
