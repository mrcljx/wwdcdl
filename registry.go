package main

import (
	"errors"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"sort"
)

type SessionResolver func(event *Event) []*Session

type Event struct {
	Id string
	Name string
	Resolver SessionResolver
}

var (
	eventId = kingpin.Arg("event", "The event to download videos/slides from").Required().String()
	events map[string]*Event
)

func RegisterEvent(event *Event) (err error) {
	if event.Resolver == nil {
		err = errors.New("Event doesn't have a Resolver")
	} else {
		events[event.Id] = event
	}

	return
}

func init() {
	events = make(map[string]*Event)
}

func EventNames() (names []string) {
	names = make([]string, 0)

	for name, _ := range events {
		names = append(names, name)
	}

	sort.Strings(names)
	return
}

func FindSessions() []*Session {
	if *eventId == "all" {
		return AllSessions()
	}

	event, ok := events[*eventId]

	if !ok {
		log.Printf("Unknown event '%s'.\n\n", *eventId)
		kingpin.Usage()
		os.Exit(1)
	}

	return event.Resolver(event)
}

func AllSessions() []*Session {
	sessions := make([]*Session, 0)

	for _, event := range events {
		sessions = append(sessions, event.Resolver(event)...)
	}

	return sessions
}
