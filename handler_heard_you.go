package main

import (
	"github.com/sromku/go-gitter"
)

type HeardYou struct{}

func (*HeardYou) Description() string {
	return "Respond blindly to each incoming message"
}

func (*HeardYou) Commands() []string {
	return []string{}
}

func (*HeardYou) Handle(msg Message, API *gitter.Gitter) {
	API.SendMessage(msg.Room.ID, "I heard you")
}
