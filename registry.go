package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"flag"
)

type SessionResolver func() []*Session

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
		fmt.Fprintf(os.Stderr, "Usage:\n\n  %s [options] event\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nEvents:\n\n")
		for _, eventName := range EventNames() {
			fmt.Fprintf(os.Stderr, "  %s (%s)\n", eventName, events[eventName].Name)
		}
		fmt.Fprintf(os.Stderr, "\nOptions:\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nNotes:\n\n")
		fmt.Fprintf(os.Stderr, "CasperJS (http://casperjs.org/) is required for authentication.\n")
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

	event, ok := events[eventId]

	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown event '%s'.\n\n", eventId)
		flag.Usage()
		os.Exit(1)
	}

	return event.Resolver()
}
