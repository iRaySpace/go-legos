package main

import "fmt"

func main() {
	lorem := "Сайт рыбатекст поможет дизайнеру"
	for _, ch := range lorem {
		fmt.Println(string(ch))
	}
}
