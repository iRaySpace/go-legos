package main

import (
	"fmt"
	"time"
)

func main() {
	timeInHours, _ := time.Parse("15", "18")
	timeInMinutes, _ := time.Parse("04", "30")

	fmt.Println(timeInHours)
	fmt.Println(timeInMinutes)
}
