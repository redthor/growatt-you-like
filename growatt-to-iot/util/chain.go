package util

import "log"

type Chain struct {
	fns   []func(message []byte)
}

func NewChain() Chain {
	return Chain{}
}

func (c *Chain) AddFn(fn func(message []byte)) (func(message []byte)){
	c.fns = append(c.fns, fn)

	log.Printf("Chain length %d.", len(c.fns))

	return c.notify
}

func (c *Chain) notify(message []byte) {
	for _, fn := range c.fns {
		fn(message)
	}
}