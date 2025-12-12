package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	data [][]string
)

func doOp(sym string, l []int) int {
	var col int
	switch sym {
	case "+":
		for _, v := range l {
			col += v
		}
	case "*":
		col = 1
		for _, v := range l {
			col *= v
		}
	default:
		panic("unknown")
	}
	return col
}

func main() {
	d, _ := os.ReadFile("input")
	lines := strings.Split(string(d), "\n")
	data = make([][]string, len(lines))
	var res1 int
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		data[i] = make([]string, len(fields))
		copy(data[i], fields)
	}

	for x := 0; x < len(data[0]); x++ {
		var l []int
		for y := 0; y < len(data)-1; y++ {
			xI, err := strconv.Atoi(data[y][x])
			if err != nil {
				panic(err)
			}
			l = append(l, xI)
		}

		res1 += doOp(data[len(data)-1][x], l)
	}
	fmt.Println("part 1", res1)
	fmt.Println("-----")
	var part2 int
	var data2 [][]byte
	data2 = make([][]byte, len(lines))
	for i, line := range lines {
		data2[i] = make([]byte, len(line))
		copy(data2[i], []byte(line))
	}
	var col []int
	for x := len(data2[0]) - 1; x >= 0; x-- {
		var lS []string
		for y := 0; y < len(data2)-1; y++ {
			s := strings.TrimSpace(string(data2[y][x]))
			if s == "" {
				continue
			}
			lS = append(lS, s)
		}
		if len(lS) > 0 {
			vS := strings.Join(lS, "")
			vI, err := strconv.Atoi(vS)
			if err != nil {
				panic(err)
			}
			col = append(col, vI)
		}
		op := data2[len(data2)-1][x]
		if op == '+' || op == '*' {
			fmt.Println(col, string(op))
			part2 += doOp(string(op), col)
			col = make([]int, 0)
		}
	}
	fmt.Println("part2:", part2)

}
