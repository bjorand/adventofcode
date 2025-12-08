package main

import (
	"fmt"
	"os"
)

type Scanner struct {
	data       []byte
	currentPos int
}

func NewScanner(data []byte) *Scanner {
	return &Scanner{data: data}
}

func (s *Scanner) isAtEnd() bool {
	return s.currentPos >= (len(s.data))
}

func (s *Scanner) advance() byte {
	s.currentPos++
	return s.data[s.currentPos-1]
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.data[s.currentPos]
}

var mulFn = []byte("mul(")

func (s *Scanner) mul(pos int) {
	if !s.isAtEnd() {
		c := s.advance()
		if c == mulFn[pos] {
			s.mul(pos + 1)
		}

	}

}

func main() {
	d, _ := os.ReadFile("input.sample")
	s := NewScanner(d)
	var mul int
	for !s.isAtEnd() {
		c := s.advance()
		switch c {
		case 'm':
			s.mul(1)
		}

	}
	fmt.Println("mul:", mul)

}
