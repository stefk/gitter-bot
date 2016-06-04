package main

import (
	"github.com/sromku/go-gitter"
)

type Handler interface {
	Init(*Context)
	Description() string
	Commands() []string
	HandleMessage(gitter.Room, gitter.Message)
}
