package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting TCP server...")
	server, _ := net.Listen("tcp", ":8080")
	defer server.Close()

	for {
		client, _ := server.Accept()
		fmt.Println("New client connected!")
		scanner := bufio.NewScanner(client)
		var read string
		for scanner.Scan() {
			read = scanner.Text()
			fmt.Printf("Was read from connection: %d bytes\n", len([]byte(read)))
			fmt.Println(read)
		}
		err := scanner.Err()
		if err != nil {
			log.Println(err)
		}
		client.Close()
	}
}
