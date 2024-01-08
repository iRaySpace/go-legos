package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"fmt"
	"strconv"
	"sort"
)

func First() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	total := 0
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "x")
		dimensions := []int{parseStrToInt(data[0]), parseStrToInt(data[1]), parseStrToInt(data[2])}
		sort.Ints(dimensions)
		paper := (dimensions[0] * dimensions[1] + dimensions[1] * dimensions[2] + dimensions[0] * dimensions[2]) * 2
		slack := dimensions[0] * dimensions[1]
		total = total + paper + slack
	}
	fmt.Println(total)
}

func Second() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	total := 0
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "x")
		dimensions := []int{parseStrToInt(data[0]), parseStrToInt(data[1]), parseStrToInt(data[2])}
		sort.Ints(dimensions)
		present := (dimensions[0] + dimensions[1]) * 2
		bow := dimensions[0] * dimensions[1] * dimensions[2]
		total = total + present + bow
	}
	fmt.Println(total)
}

func parseStrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func main() {
	First()
	Second()
}