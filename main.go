package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
)

var videos bool
var slides bool
var list bool
var authenticator *Authenticator

func init() {
	log.SetFlags(0)

	flag.BoolVar(&videos, "videos", true, "Download videos")
	flag.BoolVar(&slides, "slides", true, "Download slides/PDFs")
	flag.BoolVar(&list, "list", false, "Only list sessions")
	// TODO: flag for specific sessions
}

func main() {
	flag.Parse()

	authenticator = NewAuthenticator()
	sessions := SessionList(FindSessions())
	sort.Sort(sessions)

	log.Printf("Found %d sessions.\n", len(sessions))

	for _, session := range sessions {
		if (list) {
			fmt.Println(session.String())
			continue
		}

		if (videos) {
			DownloadVideo(session)
		}

		if (slides) {
			DownloadSlides(session)
		}
	}
}
