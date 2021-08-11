package handler

import (
	"log"
)

type GrowattProxy struct {
}

func NewGrowattProxy() *GrowattProxy {
	return &GrowattProxy{
	}
}

func (g *GrowattProxy) Register() func(message []byte) {
	log.Println("Set up Growatt Proxy.")

	return g.send
}

func (m *GrowattProxy) send(message []byte) {
	log.Printf("Proxy published %d bytes.\n", len(message))
}