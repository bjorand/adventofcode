package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hasRepetition(s string) bool {
	fmt.Println("->", s)
	if len(s) < 2 {
		return false
	}
	for i := 0; i < len(s)/2; i++ {
		prefix := s[:i+1]
		isRepeated := false
		for k := len(prefix); k < len(s); k += len(prefix) {
			if len(s)%len(prefix) > 0 {
				continue
			}
			check := s[k : len(prefix)+k]
			isRepeated = check == prefix
			if !isRepeated {
				break
			}
		}
		if isRepeated {
			return true
		}
	}
	return false
}

func main() {
	d, _ := os.ReadFile("input")
	invalidSum, invalidSum2 := 0, 0
	for _, line := range strings.Split(string(d), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		seqs := strings.Split(line, ",")
		for _, seq := range seqs {
			ii := strings.Split(seq, "-")
			if len(ii) < 2 {
				continue
			}

			i1, err := strconv.Atoi(ii[0])
			if err != nil {
				panic(err)
			}
			i2, err := strconv.Atoi(ii[1])
			if err != nil {
				panic(err)
			}
			for i := i1; i <= i2; i++ {

				s := fmt.Sprintf("%d", i)
				if len(s)%2 > 0 {
					if hasRepetition(fmt.Sprintf("%d", i)) {
						fmt.Println("R: ", i)
						invalidSum2 += i
					}
					continue
				}
				l := len(s) / 2
				s1 := string(s[0:l])
				s2 := string(s[l:])
				fmt.Println(i)
				inv, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				if s1 == s2 {
					fmt.Println("->", s)

					invalidSum += inv
				} else {
					if hasRepetition(fmt.Sprintf("%d", i)) {
						fmt.Println("R: ", i)
						invalidSum2 += i
					}
				}
			}
		}
	}
	fmt.Println(invalidSum)
	fmt.Println(invalidSum + invalidSum2)

}
