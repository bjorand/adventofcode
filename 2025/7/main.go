package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var data [][]byte
var beamLayer [][]byte
var part2 int

func dup(a [][]byte) [][]byte {
	var z [][]byte
	z = make([][]byte, len(a))
	for i := range a {
		z[i] = make([]byte, len(a[i]))
		copy(z[i], a[i])
	}
	return z
}

func draw(d [][]byte) {
	for y := range d {
		for x := range d[y] {
			fmt.Print(string(d[y][x]))
		}
		fmt.Println()
	}
}

func walk(d [][]byte, yStart int) {
	var branched bool
	for y := yStart; y < len(d); y++ {
		for x := 0; x < len(d[y]); x++ {
			c := d[y][x]
			switch c {
			case 'S':
				d[y+1][x] = '|'
			case '^':
				if d[y-1][x] == '|' {
					dL := dup(d[y:])
					dR := dup(d[y:])
					dL[0][x-1] = '|'
					walk(dL, 1)
					dR[0][x+1] = '|'
					walk(dR, 1)
					branched = true
				}
			case '.':
				if y > 0 && d[y-1][x] == '|' {
					d[y][x] = '|'
				}
			}
		}
	}
	if !branched {
		part2++
	}

}
func weightCompute(d [][]byte, yStart int) {
	w := make(map[int]int)
	for y := yStart; y < len(d); y++ {

		for x := 0; x < len(d[y]); x++ {
			c := d[y][x]
			switch c {
			case 'S':
				w[x] = 1
			case '^':
				w[x-1] += w[x]
				w[x+1] += w[x]
				w[x] = 0
			}
		}
		fmt.Println("line", y)

	}
	for i := range w {
		part2 += w[i]
	}

}

func main() {
	d, _ := os.ReadFile("input")
	lines := strings.Split(string(d), "\n")
	data = make([][]byte, 0)
	beamLayer = make([][]byte, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		beamLayer[i] = make([]byte, len(line))
		useless := true
		for k := range line {
			if line[k] == '^' || line[k] == 'S' {
				useless = false
			}
		}
		if !useless || i == 1 {
			data = append(data, []byte(line))
		}
	}
	draw(data)

	var part1 int
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			c := data[y][x]

			switch c {
			case 'S':
				beamLayer[y+1][x] = '|'
			case '^':
				if beamLayer[y-1][x] == '|' {
					beamLayer[y][x-1] = '|'
					beamLayer[y][x+1] = '|'
					part1++

				}
			case '.':
				if y > 0 && beamLayer[y-1][x] == '|' {
					beamLayer[y][x] = '|'
				}
			}
		}
	}
	fmt.Println("part1:", part1)
	go func() {
		t := time.NewTicker(time.Second)
		var prev int
		for {
			<-t.C

			fmt.Println(part2-prev, "/s")
			prev = part2
		}
	}()
	// brute force approach of would take 114 days to compute at 1M branch/s /o\
	// walk(data, 0)
	weightCompute(data, 0)
	fmt.Println("part2:", part2)
}
