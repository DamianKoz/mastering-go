package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
