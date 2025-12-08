package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findSimilarityScore(i int, l []int) int {
	var found int
	for _, v := range l {
		if v == i {
			found += 1
		}
	}
	return found * i
}

func main() {
	var l1, l2 []int
	d, _ := os.ReadFile("input")
	for _, line := range strings.Split(string(d), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		str := strings.Fields(line)
		i1, err := strconv.Atoi(str[0])
		if err != nil {
			panic(err)
		}
		i2, err := strconv.Atoi(str[1])
		if err != nil {
			panic(err)
		}
		l1 = append(l1, i1)
		l2 = append(l2, i2)
	}
	sort.Ints(l1)
	sort.Ints(l2)
	var distance int
	var similarityScore int
	for i := 0; i < len(l1); i++ {
		v1 := l1[i]
		similarityScore += findSimilarityScore(v1, l2)
		v2 := l2[i]
		distance += max(v1, v2) - min(v1, v2)
	}
	fmt.Println("part 1:", distance)
	fmt.Println("part 2:", similarityScore)

}
