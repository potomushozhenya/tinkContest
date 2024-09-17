package main

import (
	"fmt"
	"math"
)

func divisorCount(x int) int {
	c := 0
	limit := int(math.Sqrt(float64(x))) + 1
	for i := 1; i < limit; i++ {
		if x%i == 0 {
			if i*i != x {
				c += 2
			} else {
				c += 1
			}

		}
	}
	return c
}
func main() {
	// Подходят только полные квадраты
	seive := make([]bool, 400)

	limit := int(math.Sqrt(float64(400))) + 1

	for i := 2; i < limit; i++ {
		if !seive[i] {
			for j := i * i; j < 400; j += i {
				seive[j] = true
			}
		}
	}
	var l, r int
	fmt.Scanln(&l, &r)
	count := 0
	curr := int(math.Ceil(math.Sqrt(float64(l))))
	if curr == 1 {
		curr++
	}
	for curr*curr <= r {
		if !seive[divisorCount(curr*curr)] {
			count++
		}
		curr++
	}
	fmt.Println(count)
}
