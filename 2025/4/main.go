package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	m              [][]byte
	reachableRolls int
)

func collect(removeRoll bool, yStart int) {
	for y := yStart; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			switch m[y][x] {
			case '.':
				fmt.Print(".")
			case '@':
				var toScan []byte
				// add top
				if y > 0 {
					toScan = append(toScan, m[y-1][max(0, x-1):min(x+2, len(m[y]))]...)
				}
				// add bottom
				if y < len(m)-1 {
					toScan = append(toScan, m[y+1][max(0, x-1):min(x+2, len(m[y]))]...)
				}
				// add left
				if x > 0 {
					toScan = append(toScan, m[y][x-1])
				}
				// add right
				if x < len(m[y])-1 {
					toScan = append(toScan, m[y][x+1])
				}
				rolls := 0
				for c := range toScan {
					if toScan[c] == '@' {
						rolls++
					}
				}
				fmt.Print("@")
				if rolls < 4 {
					reachableRolls++
					if removeRoll {
						m[y][x] = '.'
						collect(true, max(0, y-1))
					}
				}
			}

		}
		fmt.Print("\n")
	}
}

func main() {
	d, _ := os.ReadFile("input")
	lines := strings.Split(string(d), "\n")
	m = make([][]byte, len(lines))

	y := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		m[y] = make([]byte, len(line))
		x := 0
		for i := range line {
			m[y][x] = line[i]
			x++
		}
		y++
	}

	collect(false, 0)
	fmt.Println(reachableRolls)
	reachableRolls = 0

	collect(true, 0)
	fmt.Println(reachableRolls)

}
