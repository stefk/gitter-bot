package main

import (
	"github.com/sromku/go-gitter"
)

type Listener func (data interface{})

type Context struct {
	BotUser *gitter.User
	BotRooms []gitter.Room
	API *gitter.Gitter
	listeners map[string][]Listener
}

func (c *Context) SendMessage(roomID string, text string) error {
	return c.API.SendMessage(roomID, text)
}

func (c *Context) On(event string, listener Listener) {
	if c.listeners == nil {
		c.listeners = make(map[string][]Listener)
	}

	_, ok := c.listeners[event]

	if !ok {
		c.listeners[event] = make([]Listener, 0)
	}

	c.listeners[event] = append(c.listeners[event], listener)
}

func (c *Context) Emit(event string, data interface{}) {
	if c.listeners == nil {
		c.listeners = make(map[string][]Listener)
	}

	_, ok := c.listeners[event]

	if ok {
		for _, listener := range c.listeners[event] {
			go listener(data)
		}
	}
}
