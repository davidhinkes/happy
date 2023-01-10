// Binary happy evaluates integers for happy numbers.
package main

import (
	"flag"
	"fmt"
)

var (
	searchSize = flag.Int("i", 100, "find happy numbers from [0,i)")
)

func main() {
	flag.Parse()
	for i := 0; i < *searchSize; i++ {
		h, l := isHappy(i)
		if !h {
			continue
		}
		fmt.Printf("isHappy(%v): %v, %v\n", i, h, l)
	}
}

// next is recursive means for gennerating the next potential happy number
// given x.
// For a given x, we expect log_10(x) calls to nextHappy, so recursive is OK.
func next(x int) int {
	if x == 0 {
		return x
	}
	// get the lowest digit of x
	d := x % 10
	// remove the lowest digit and grap the next of this.
	// ex: x=123
	//     return 3*3 + next(12)
	return d*d + next(x/10)
}

// isHappy determines if x is a happy number (bool), and also returns the sequence of next-happies.
func isHappy(x int) (bool, []int) {
	m := map[int]struct{}{}
	return isHappyR([]int{x}, m)
}

func isHappyR(xs []int, m map[int]struct{}) (bool, []int) {
	x := xs[len(xs)-1] // assume xs not empty
	_, dup := m[x]
	m[x] = struct{}{}
	xs = append(xs, next(x))
	if dup {
		return false, xs
	}
	if x == 1 {
		return true, xs
	}
	return isHappyR(xs, m)
}
