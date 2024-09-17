package main

import "fmt"

func main() {
	var n, currH int
	h := 0
	isOk := true
	fmt.Scan(&n)
	dHeight := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&currH)
		if currH != -1 {
			if h > currH {
				isOk = false
				break
			}
			dHeight[i] = currH - h
			h += currH - h
		} else {
			dHeight[i] = 1
			h += 1
		}
	}
	if isOk {
		fmt.Println("YES")
		for i := 0; i < n; i++ {
			fmt.Print(dHeight[i], " ")
		}
	} else {
		fmt.Println("NO")
	}
}
