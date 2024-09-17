package main

import (
	"bufio"
	"fmt"
	"os"
)

func arrClean(arr []int) {
	for i := 0; i < 256; i++ {
		if arr[i] > 1 {
			arr[i] = 1
		}
	}
}
func main() {
	left := -1
	right := -1
	var k int
	var isSol bool
	i := 0
	reader := bufio.NewReader(os.Stdin)
	seq, _ := reader.ReadString('\n')
	request, _ := reader.ReadString('\n')
	fmt.Scan(&k)
	var chars [256]int
	for request[i] != '\n' {
		chars[int(request[i])] = 1
		i++
	}
	i = 0
	for seq[i] != '\n' {
		isSol = true
		j := i
		for seq[j] != '\n' && j < i+k {
			ind := int(seq[j])
			if chars[ind] >= 1 {
				chars[ind]++
			}
			j++
		}
		for k := 0; k < 256; k++ {
			if chars[k] == 1 {
				isSol = false
				break
			}
		}
		arrClean(chars[:])
		if isSol {
			left = i
			right = j
		}
		i++
	}
	if left != -1 {
		fmt.Print(seq[left:right])
	} else {
		fmt.Print(-1)
	}
	fmt.Print("\n")
}
