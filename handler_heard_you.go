package main

import (
	"github.com/sromku/go-gitter"
)

type HeardYou struct {
	Context *Context
}

func (h *HeardYou) Init(context *Context) {
	h.Context = context
}

func (*HeardYou) Description() string {
	return "Respond blindly to each incoming message"
}

func (*HeardYou) Commands() []string {
	return []string{}
}

func (h *HeardYou) HandleMessage(room gitter.Room, msg gitter.Message) {
	h.Context.SendMessage(room.ID, "I heard you")

	if msg.Text == "secret!" {
		h.Context.Emit("foo", Bot{})
	}
}
