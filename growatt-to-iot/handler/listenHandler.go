package handler

import (
	"fmt"
	"log"
	"net"
)

type ListenHandler struct {
	ip         net.IP
	port       int
	packetSize int
	onMessage  func(message []byte)
}

func NewListenHandler(ip net.IP, port int, packetSize int, onMessage func(message []byte)) *ListenHandler {
	return &ListenHandler{
		ip:         ip,
		port:       port,
		packetSize: packetSize,
		onMessage:  onMessage,
	}
}

func (l *ListenHandler) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", l.ip.String(), l.port))
	if err != nil {
		log.Fatalf("Cannot listen: %v", err.Error())
	}
	defer listener.Close()

	log.Println("Started listening.")

	message := make([]byte, l.packetSize)

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Accept connecton failed: %v", err.Error())
	}

	for {
		n, err := conn.Read(message)
		if err != nil {
			log.Fatalf("Received an error reading from the socket: %v", err.Error())
		}
		fmt.Printf("Read %d bytes\n", n)
		l.onMessage(message[:n])
	}
}
