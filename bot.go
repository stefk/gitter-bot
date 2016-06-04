package main

import (
	"github.com/sromku/go-gitter"
	"log"
	"sync"
)

type Bot struct {
	API      *gitter.Gitter
	user     *gitter.User
	rooms    []gitter.Room
	handlers []Handler
}

type Event struct {
	Room gitter.Room
	Data interface{}
}

type Message struct {
	Room    gitter.Room
	Message gitter.Message
}

type Handler interface {
	Description() string
	Commands() []string
	Handle(Message, *gitter.Gitter)
}

type stream struct {
	Room   gitter.Room
	Stream *gitter.Stream
}

func NewBot(token string) *Bot {
	API := gitter.New(token)
	user, err := API.GetUser()

	if err != nil {
		log.Println("Cannot get user from gitter API")
		log.Fatal(err)
	}

	log.Printf("Connected to gitter API as %s\n", user.Username)

	rooms, err := API.GetRooms()

	if err != nil {
		log.Println("Cannot get rooms from gitter API")
		log.Fatal(err)
	}

	log.Printf("Found %d available room(s)\n", len(rooms))

	handlers := make([]Handler, 0)

	return &Bot{API, user, rooms, handlers}
}

func (bot *Bot) AddHandler(handler Handler) {
	bot.handlers = append(bot.handlers, handler)
}

// Start streaming available rooms. Handlers will be triggered on each
// incoming message
func (bot *Bot) Listen() {
	streams := make([]stream, 0)

	for _, room := range bot.rooms {
		log.Printf("Streaming room %s\n", room.URL)
		roomStream := bot.API.Stream(room.ID)
		streams = append(streams, stream{Room: room, Stream: roomStream})
		go bot.API.Listen(roomStream)
	}

	c := merge(streams)

	for {
		event := <-c
		switch ev := event.Data.(type) {
		case *gitter.MessageReceived:
			if ev.Message.From.Username != bot.user.Username {
				log.Printf("Received message from %s\n", event.Room.URL)
				msg := Message{Room: event.Room, Message: ev.Message}

				for _, handler := range bot.handlers {
					handler.Handle(msg, bot.API)
				}
			}
		case *gitter.GitterConnectionClosed:
			// connection was closed
		}
	}
}

// Merge event channels from different rooms into a single channel
// (see https://blog.golang.org/pipelines)
func merge(streams []stream) chan Event {
	var wg sync.WaitGroup
	out := make(chan Event)

	output := func(s stream) {
		for n := range s.Stream.Event {
			out <- Event{Room: s.Room, Data: n.Data}
		}
		wg.Done()
	}

	wg.Add(len(streams))

	for _, s := range streams {
		go output(s)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
