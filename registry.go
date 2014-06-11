package main

import (
	"errors"
	"flag"
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

var events map[string]*Event

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

	flag.Usage = func() {
		log.Printf("Usage:\n\n  %s [options] event\n", os.Args[0])
		log.Println("\nEvents:\n")
		log.Println("  all")
		for _, eventName := range EventNames() {
			log.Printf("  %s (%s)\n", eventName, events[eventName].Name)
		}
		log.Println("\nOptions:\n")
		flag.PrintDefaults()
		log.Println("\nNotes:\n")
		log.Println("CasperJS (http://casperjs.org/) is required for authentication.")
	}
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
	eventId := flag.Arg(0)

	if len(eventId) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if eventId == "all" {
		return AllSessions()
	}

	event, ok := events[eventId]

	if !ok {
		log.Printf("Unknown event '%s'.\n\n", eventId)
		flag.Usage()
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
