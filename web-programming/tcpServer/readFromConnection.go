package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"time"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err!= nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
			if err!= nil {
				log.Fatalln(err)
			}
			go handle(conn)
		}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(5* time.Second))
	if err!= nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	defer conn.Close()
	fmt.Println("**This code ran.**")
}