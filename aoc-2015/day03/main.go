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
	houses := 1 // Already visited in his location?!?
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
	houses = houses + len(m)
	fmt.Println(houses)
}

func Second() {

}

func main() {
	First()
	Second()
}
