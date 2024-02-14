package main

import (
	"log"

	"github.com/irayspace/go-legos/learn-folder-structure/internal"
	"github.com/irayspace/go-legos/learn-folder-structure/pkg/math"
)

func main() {
	log.Println("Starting application...")
	internal.InitDB()
	log.Println("Ivan's number is", math.GetIvanNo(9))
}
