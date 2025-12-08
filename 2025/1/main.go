package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dial := 50
	pointToZero := 0
	pointToZero2 := 0
	d, _ := os.ReadFile("input")
	for _, line := range strings.Split(string(d), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		dir := string(line[0])

		incr, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			panic(err)
		}
		switch dir {
		case "L":
			for i := dial; i > dial-incr; i-- {
				if i%100 == 0 || i == 0 {
					pointToZero2++
				}
			}
			dial = dial - incr%100
			if dial < 0 {
				dial = 100 + dial
			}
		case "R":
			for i := dial; i < dial+incr; i++ {
				if i%100 == 0 || i == 0 {
					pointToZero2++
				}
			}
			dial = (dial + incr%100) % 100

		default:
			panic("Unknown dir")
		}
		if dial == 0 {
			pointToZero++
		}

	}
	fmt.Println(pointToZero)
	fmt.Println(pointToZero2)

}
