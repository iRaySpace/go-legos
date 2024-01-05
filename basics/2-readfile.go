package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("../res/lorem.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
