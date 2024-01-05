package main

import "fmt"

func main() {
	nos := []int{1, 2, 3, 4}
	nos = append(nos, 5)
	fmt.Println(nos)
}
