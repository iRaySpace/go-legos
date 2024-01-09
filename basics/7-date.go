package main

import (
	"fmt"
	"time"
)

func main() {
	_, month, day := time.Now().Date()
	if int(month) == 1 && day == 1 {
		fmt.Println("Happy New Year!")
	}
	if int(month) == 1 && day == 7 {
		fmt.Println("С Рождеством!")
	}
	if int(month) == 1 && day == 9 {
		fmt.Println("Today is the Feast of the Black Nazarene")
	}
}
