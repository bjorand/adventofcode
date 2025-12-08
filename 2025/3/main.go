package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	d, _ := os.ReadFile("input")
	total := 0
	total12 := 0
	for _, line := range strings.Split(string(d), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		max := 0
		var batteries []int
		for i := 0; i < len(line); i++ {
			bat, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)

			}
			batteries = append(batteries, bat)
		}

		for dI, dV := range batteries {
			if dI >= len(batteries)-1 {
				break
			}
			for i := dI + 1; i < len(batteries); i++ {
				uV := batteries[i]
				cur := dV*10 + uV
				if cur > max {
					max = cur
				}

			}

		}
		total += max

		pos := 0
		var res = make([]int, 12)
		for i := range res {
			res[i] = -1
		}
		for i := 0; i < 12; i++ {
			fmt.Println(pos)
			for k := pos; k < len(batteries)-11+i; k++ {
				if batteries[k] > res[i] {
					res[i] = batteries[k]
					pos = k + 1
				}

			}

		}
		var resS string
		for i := range res {
			resS += fmt.Sprintf("%d", res[i])
		}
		sum, err := strconv.Atoi(resS)
		if err != nil {
			panic(err)
		}
		total12 += sum
		fmt.Println("-")
	}
	fmt.Println(total)
	fmt.Println(total12)
}
