package main

import (
	"fmt"
	"os"
	"sort"
	"flag"
)

type SessionResolver func() []*Session

var events map[string]SessionResolver

func init() {
	events = make(map[string]SessionResolver)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n\n  %s [options] event\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nEvents:\n\n")
		for _, eventName := range EventNames() {
			fmt.Fprintf(os.Stderr, "  %s\n", eventName)
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
	event := flag.Arg(0)

	if len(event) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	resolver, ok := events[event]

	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown event '%s'.\n\n", event)
		flag.Usage()
		os.Exit(1)
	}

	return resolver()
}
