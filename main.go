package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"os"
)

const DEFAULT_PORT = "8007"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	server, err := net.Listen("tcp", ":"+port)
	if server == nil {
		log.Fatalln(err)
	}
	defer server.Close()

	log.Printf("Listening on :%s", port)

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Printf("localAddr: %v <-> RemoteAddr: %v - CONNECTED", conn.LocalAddr(), conn.RemoteAddr())
		go echo(conn)
	}
}

func echo(conn net.Conn) {
	b := bufio.NewReader(conn)
	bytesRead := 0
	for {
		line, err := b.ReadBytes('\n')
		bytesRead += len(line)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Printf("localAddr: %v <-> RemoteAddr: %v - DISCONNECTED (bytes read/written: %d)", conn.LocalAddr(), conn.RemoteAddr(), bytesRead)
				conn.Close()
				break
			}
			break
		}
		conn.Write(line)
	}
}
