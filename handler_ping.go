package main

import (
	"github.com/sromku/go-gitter"
	"log"
)

type Ping struct {
	Context *Context
}

func (h *Ping) Init(context *Context) {
	h.Context = context
	h.Context.On("foo", func(data interface{}) {
		log.Println("Received foo event")
	})
}

func (*Ping) Description() string {
	return "Respond to ping messages"
}

func (*Ping) Commands() []string {
	return []string{"ping - Respond with pong"}
}

func (h *Ping) HandleMessage(room gitter.Room, msg gitter.Message) {
	if msg.Text == "ping" {
		h.Context.SendMessage(room.ID, "Pong")
	}
}
