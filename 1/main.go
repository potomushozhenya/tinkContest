package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func genSeq(a int, b int, arr *[]int) {
	for i := a; i <= b; i++ {
		*arr = append(*arr, i)
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	l := 0
	r := 1
	left := 0
	right := 0
	res := []int{}
	isSegment := false
	for line[r-1] != '\n' {
		if line[r] == '-' {
			if l != r-1 {
				left, _ = strconv.Atoi(line[l:r])
			} else {
				left = int(line[l]) - '0'
			}
			l = r + 1
			isSegment = true
		} else {
			if line[r] == '\n' || line[r] == ',' {
				if l != r-1 {
					right, _ = strconv.Atoi(line[l:r])
				} else {
					right = int(line[l]) - '0'
				}
				l = r + 1
				if isSegment == true {
					genSeq(left, right, &res)
					isSegment = false
				} else {
					res = append(res, right)
				}

			}
		}
		r++
	}
	for i := 0; i < len(res)-1; i++ {
		fmt.Printf("%d", res[i])
		fmt.Print(" ")
	}
	fmt.Print(res[len(res)-1], "\n")
}
