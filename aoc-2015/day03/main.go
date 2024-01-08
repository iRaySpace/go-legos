package main

import (
	"fmt"
	"log"
	"os"
)

func First() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[string]bool)
	m["0,0"] = true
	var x, y int
	for _, ch := range data {
		switch ch {
		case '^':
			y = y + 1
		case 'v':
			y = y - 1
		case '>':
			x = x + 1
		case '<':
			x = x - 1
		}
		m[fmt.Sprintf("%d,%d", x, y)] = true
	}
	fmt.Println(len(m))
}

func Second() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[string]bool)
	m["0,0"] = true
	var x, y int
	var rx, ry int
	for i, ch := range data {
		robot := i%2 != 0
		switch ch {
		case '^':
			if robot {
				ry = ry + 1
			} else {
				y = y + 1
			}
		case 'v':
			if robot {
				ry = ry - 1
			} else {
				y = y - 1
			}
		case '>':
			if robot {
				rx = rx + 1
			} else {
				x = x + 1
			}
		case '<':
			if robot {
				rx = rx - 1
			} else {
				x = x - 1
			}
		}
		if robot {
			m[fmt.Sprintf("%d,%d", rx, ry)] = true
		} else {
			m[fmt.Sprintf("%d,%d", x, y)] = true
		}
	}
	fmt.Println(len(m))
}

func main() {
	First()
	Second()
}
