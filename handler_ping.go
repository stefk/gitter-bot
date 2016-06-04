package main

import (
	"github.com/sromku/go-gitter"
)

type Ping struct{}

func (*Ping) Description() string {
	return "Respond to ping messages"
}

func (*Ping) Commands() []string {
	return []string{"ping - Respond with pong"}
}

func (*Ping) Handle(msg Message, API *gitter.Gitter) {
	if msg.Message.Text == "ping" {
		API.SendMessage(msg.Room.ID, "Pong")
	}
}
