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
	floor := 0
	for _, ch := range data {
		if ch == '(' {
			floor = floor + 1
		} else if ch == ')' {
			floor = floor - 1
		}
	}
	fmt.Println(floor)
}

func Second() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	floor := 0
	for i, ch := range data {
		if ch == '(' {
			floor = floor + 1
		} else if ch == ')' {
			floor = floor - 1
		}
		if floor == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}

func main() {
	First()
	Second()
}
