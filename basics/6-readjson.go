package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("../res/lorem.json")
	if err != nil {
		panic(err)
	}
	var jsonMap map[string]interface{}
	err = json.Unmarshal(data, &jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonMap["id"], jsonMap["message"])
}