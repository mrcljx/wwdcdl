package main

import (
	"fmt"
	"flag"
)

var videos bool
var slides bool
var list bool
var authenticator *Authenticator

func init() {
	flag.BoolVar(&videos, "videos", true, "Download videos")
	flag.BoolVar(&slides, "slides", true, "Download slides/PDFs")
	flag.BoolVar(&list, "list", false, "Only list sessions")
	// TODO: flag for specific sessions
}

func main() {
	flag.Parse()

	authenticator = NewAuthenticator()
	sessions := FindSessions()

	fmt.Printf("Found %d sessions.\n", len(sessions))

	for _, session := range sessions {
		if (list) {
			fmt.Printf("%s", session)
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
