package main

import (
	"net"
	"flag"
	"fmt"
	"log"
)

func main() {
	portPtr := flag.Int("port", 8080, "the port to listen on")
	flag.Parse()
	addr := fmt.Sprintf(":%d", *portPtr)

	fmt.Printf("Starting web server... listening on address %v\n", addr)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection%v\n", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	log.Printf("Connection opened...\n")

	requestBytes := make([]byte, 1024)
	bytesRead, err := conn.Read(requestBytes)
	if err != nil {
		log.Printf("Error reading from the connection:%v\n", err)

		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing the connection: %v\n", err)
		}
	}()
	request := string(requestBytes)

	log.Printf("Read %d bytes:\n%v\n-----EOF-----\n", bytesRead, request)

	html := "<html><body><blink>Hello world</blink></body></html>"
	contentLength := len(html)
	response := fmt.Sprintf(
`HTTP/1.0 200 OK
Content-Length: %d

%v
`, contentLength, html)

	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Error writing to the connection:%v\n", err)
	}
}
