package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

var exit = false

func main() {
	c, err := net.Dial("tcp", "localhost:2567")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	log.Println("Connected on :2567")

	go handleRead(c)

	r := bufio.NewReader(os.Stdin)
	for {
		text, _ := r.ReadBytes('\n')

		n, _ := c.Write(text)
		log.Printf("Wrote %d bytes\n", n)

		if exit {
			break
		}
	}
	// r := bufio.NewReader(os.Stdin)
	// sr := bufio.NewReader(c)
	// for {
	// 	fmt.Print("Input: ")
	// 	text, _ := r.ReadBytes('\n')

	// 	n, err := c.Write(text)
	// 	if err != nil {
	// 		log.Println("Unable to send request")
	// 	}
	// 	res, err := sr.ReadString('\n')
	// 	if err != nil {
	// 		log.Println("Unable to read response")
	// 		continue
	// 	}
	// 	log.Println(res)
	// }
}

func handleRead(c net.Conn) {
	reader := bufio.NewReader(c)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			continue
		}

		res := strings.TrimSuffix(string(line), "\n")
		log.Printf("Server: %s", res)

		if res == "win" {
			exit = true
			break
		}
	}
}
