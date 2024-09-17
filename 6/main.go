package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	type Proc struct {
		procID int
		time   uint64
		dep    []int
	}
	var n int
	fmt.Scan(&n)
	p := make([]Proc)
	for i := 0; i < n; i++ {
		var curr uint64
		fmt.Scan(&curr)
		newP := new(Proc)
		newP.time = curr
		reader := bufio.NewReader(os.Stdin)
		seq, _ := reader.ReadString('\n')
		l := 0
		r := 0
		for seq[r] != '\n' || seq[r] != '\r' {
			if seq[r] == ' ' {

			}
		}
	}
}
