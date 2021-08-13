package handler

import (
	"fmt"
	"log"
	"net"
)

var (
	// Hard coded for now
	growattServerIP = net.IPv4(47, 91, 67, 66)
	growattPort     = 5279
)

type GrowattProxy struct {
	conn net.Conn
}

func NewGrowattProxy() *GrowattProxy {
	return &GrowattProxy{}
}

func (g *GrowattProxy) Register() func(message []byte) {
	log.Println("Set up Growatt Proxy.")

	var err error
	g.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", growattServerIP.String(), growattPort))
	if err != nil {
		log.Fatalf("Error connecting to growatt server. Err = %s", err.Error())
	}

	return g.send
}

func (g *GrowattProxy) send(message []byte) {
	count, err := g.conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("Error sending to growatt server. Err = %s", err.Error())
	}
	log.Printf("Growatt proxy sent [%d] bytes.", count)
}
