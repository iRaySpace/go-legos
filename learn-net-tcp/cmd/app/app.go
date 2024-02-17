package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":2567")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	log.Println("Running on :2567")

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("Problem accepting incoming connection from '" + c.RemoteAddr().String() + "'")
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	n, err := c.Write([]byte("Welcome to learn-net-tcp\n"))
	if err != nil {
		log.Println("Unable to send to '" + c.RemoteAddr().String() + "'")
	}
	log.Printf("Wrote %d bytes", n)

	reader := bufio.NewReader(c)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			c.Close()
			return
		}

		res := strings.TrimSuffix(line, "\n")
		log.Printf("Client [%s]: %s", c.RemoteAddr().String(), res)

		if res == "ivanof" {
			n, _ := c.Write([]byte("win\n"))
			log.Printf("Wrote %d bytes", n)
			c.Close()
			break
		}

		n, _ := c.Write([]byte("invalid-keyword\n"))
		log.Printf("Wrote %d bytes", n)
	}
}
