package main

import (
	"flag"
	"log"
	"net"
	"strconv"
)

func main() {
	port := flag.Int("port", 33333, "Port to accept connections on.")
	flag.Parse()

	l, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Listening to connections on port", strconv.Itoa(*port))
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panicln(err)
		}

		tcpConn, _ := conn.(*net.TCPConn)
		go handleRequest(tcpConn)
	}
}

func handleRequest(conn *net.TCPConn) {
	log.Println("Accepted new connection.")
	defer conn.Close()
	defer log.Println("Closed connection.")

	buf := make([]byte, 16192)
	for {
		size, err := conn.Read(buf)
		if err != nil {
			return
		}
		data := buf[:size]
		// log.Println("Read new data from connection", data)
		conn.Write(data)
	}
}
