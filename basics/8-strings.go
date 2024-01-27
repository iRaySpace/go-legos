package main

import (
	"fmt"
	"strings"
)

func main() {
	row := strings.Split("Ivan Ray,Altomera", ",")
	fname, lname := row[0], row[1]
	fmt.Println(lname, ",", fname)

	names := []string{"Ivan Ray", "Altomera"}
	fullName := strings.Join(names, " ")
	fmt.Println(fullName)

	simpleName := "ivan"
	fmt.Println(strings.ToUpper(simpleName))

	simpleName = "IVAN"
	fmt.Println(strings.ToLower(simpleName))
}
