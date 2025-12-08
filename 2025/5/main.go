package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ingredientsID     []int
	freshRanges       []*freshRange
	allFreshIDs       []int
	freshIDcount      int
	mergedFreshRanges []*freshRange
)

type freshRange struct {
	min int
	max int
}

func merge(orig []*freshRange) []*freshRange {
	fmt.Println("-----")
	for _, r := range orig {
		fmt.Println(r.min, "-", r.max)
	}
	var res []*freshRange
	var needRemerge bool
	for _, r := range orig {
		var exists bool
		for _, m := range res {
			if r.min >= m.min && r.max <= m.max {
				exists = true
				break
			}
			// left extends
			if r.min < m.min && r.max >= m.min {
				m.min = r.min
				exists = true
				needRemerge = true
			}
			// right extends
			if r.min <= m.max && r.max >= m.max {
				m.max = r.max
				exists = true
				needRemerge = true
			}
		}
		if !exists {
			res = append(res, r)
		}
	}
	if needRemerge {
		fmt.Println("remerge")
		res = merge(res)
	}
	return res
}

func main() {
	d, _ := os.ReadFile("input")

	for _, line := range strings.Split(string(d), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		l := strings.Split(line, "-")
		switch len(l) {
		case 1:
			id, err := strconv.Atoi(l[0])
			if err != nil {
				panic(err)
			}
			ingredientsID = append(ingredientsID, id)
		case 2:
			minS := l[0]
			maxS := l[1]
			minI, err := strconv.Atoi(minS)
			if err != nil {
				panic(err)
			}
			maxI, err := strconv.Atoi(maxS)
			if err != nil {
				panic(err)
			}
			freshRanges = append(freshRanges, &freshRange{min: minI, max: maxI})
		}
	}
	freshCount := 0
	for _, ingredientID := range ingredientsID {
		freshCount += func() int {
			for _, r := range freshRanges {
				if ingredientID >= r.min && ingredientID <= r.max {
					return 1
				}
			}
			return 0
		}()
	}
	fmt.Println(freshCount)
	mergedFreshRanges = merge(freshRanges)
	fmt.Println("before merge: ", len(freshRanges), " after:", len(mergedFreshRanges))

	for _, r := range mergedFreshRanges {
		fmt.Println(r.min, "-", r.max)
	}

	for _, r := range mergedFreshRanges {
		freshIDcount += r.max - r.min + 1

	}
	fmt.Println(freshIDcount)
}
