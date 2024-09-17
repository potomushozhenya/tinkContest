package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func timeConvert(line string) int {
	h, _ := strconv.Atoi(line[0:2])
	m, _ := strconv.Atoi(line[3:5])
	s, _ := strconv.Atoi(line[6:8])
	return h*60*60 + m*60 + s
}
func countPenalty(t int, s int, init int) int {
	if init > t {
		return (t - s - init + 24*60*60) / 60
	} else {
		return (t - s - init) / 60
	}
}
func main() {
	type Team struct {
		points      int
		penalty     int
		servPending [26]int
	}
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	init := timeConvert(line)
	limit := init + 24*60*60
	var n int
	table := map[string]*Team{}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var team, time, serv, status string
		fmt.Scan(&team, &time, &serv, &status)
		t := timeConvert(time)
		if t < limit {
			val, ok := table[team]
			if ok {
				if status != "PONG" {
					if status == "ACCESSED" {
						val.points += 1
						s, _ := strconv.Atoi(line[6:8])
						val.penalty += countPenalty(t, s, init)
						val.penalty += val.servPending[int(serv[0])-'A']
					} else {
						val.servPending[int(serv[0])-'A'] += 20
					}
				}
			} else {
				newStruct := new(Team)
				if status != "PONG" {
					if status == "ACCESSED" {
						newStruct.points = 1
						s, _ := strconv.Atoi(line[6:8])
						newStruct.penalty = countPenalty(t, s, init)
					} else {
						newStruct.points = 0
						newStruct.servPending[int(serv[0])-'A'] += 20
					}
				} else {
					newStruct.points = 0
					newStruct.penalty = 0
				}
				table[team] = newStruct
			}
		}
	}
	keys := make([]string, 0, len(table))
	for key := range table {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if table[keys[i]].points != table[keys[j]].points {
			return table[keys[i]].points > table[keys[j]].points
		} else {
			return table[keys[i]].penalty < table[keys[j]].penalty
		}
	})
	for i, k := range keys {
		fmt.Println(i+1, k, table[k].points, table[k].penalty)
	}
}
