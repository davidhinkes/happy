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
		fmt.Printf("isHappy(%v): %v, %v\n", i, h, l)
	}
}

// next is recursive means for gennerating the next potential happy number
// given x.
// For a given x, we expect log(x) calls to nextHappy, so recursive is OK.
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
	var happies []int
	m := map[int]struct{}{}
	for {
		happies = append(happies, x)
		if _, ok := m[x]; ok || x == 1 {
			break
		}
		m[x] = struct{}{}
		x = next(x)
	}
	return x == 1, happies
}
