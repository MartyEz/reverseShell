package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Read from connection stream
func read(conn *net.Conn) {
	buf := make([]byte, 4080)
	for {
		n, err := (*conn).Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		receivedCmd := string(buf[:n])
		fmt.Println(receivedCmd)
	}
}

func main() {

	// Setting up a connection listener
	fmt.Println("Waiting client connection\n")
	ln, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}

	// Waiting a connection from client
	conn, err := ln.Accept()
	fmt.Print("Get connection from :", conn.RemoteAddr(), " at : ", conn.LocalAddr())
	if err != nil {
		log.Fatal(err)
	}

	// Launch goroutine to read connection stream
	go read(&conn)

	// Infinite loop. Wait string cmd from stdin and send it to conn stream as []bytes
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "quitServer"{
			return
		}
		text = text + "\n"
		_, err := conn.Write([]byte(text))
		if err != nil {
			log.Fatal(err)
		}
	}
}
