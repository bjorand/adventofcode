package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func removeForList(pos int, l []int) []int {
	var new []int
	for i := 0; i < len(l); i++ {
		if pos == i {
			continue
		}
		new = append(new, l[i])
	}
	return new
}

func isSafe(levels []int) bool {
	var decr bool
	var incr bool
	switch {
	case levels[0] < levels[1]:
		incr = true
	case levels[0] > levels[1]:
		decr = true
	}
	for i := 1; i < len(levels); i++ {
		if levels[i-1] > levels[i] && incr {
			return false
		}
		if levels[i-1] < levels[i] && decr {
			return false

		}
		diff := max(levels[i-1], levels[i]) - min(levels[i-1], levels[i])
		if diff < 1 || diff > 3 {
			return false
		}

	}
	return true
}

func main() {
	d, _ := os.ReadFile("input")
	var safeReportsCount int
	var safeReportsCountMinusOne int
	for _, line := range strings.Split(string(d), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		levelsString := strings.Fields(line)
		var levels []int
		for _, s := range levelsString {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			levels = append(levels, i)
		}

		safe := isSafe(levels)
		if safe {
			safeReportsCount++
		} else {
			safeReportsCountMinusOne += func() int {
				for toRemoveIdx := 0; toRemoveIdx < len(levels); toRemoveIdx++ {
					if isSafe(removeForList(toRemoveIdx, levels)) {
						return 1
					}
				}
				return 0
			}()
		}

	}
	fmt.Println("fully safe levels:", safeReportsCount)
	fmt.Println("safe levels minus one:", safeReportsCountMinusOne+safeReportsCount)
}
