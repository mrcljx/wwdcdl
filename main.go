package main

import (
	"bytes"
	"fmt"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"log"
	"sort"
)

var (
  noVideos = kingpin.Flag("no-videos", "Don't download videos").Bool()
  noSlides = kingpin.Flag("no-slides", "Don't download slides/PDFs").Bool()
  list = kingpin.Flag("list", "Only show list of sessions").Short('l').Bool()
)

var authenticator *Authenticator

func init() {
	log.SetFlags(0)
}

func configure() {
	buffer := bytes.NewBufferString(kingpin.DefaultUsageTemplate)

	fmt.Fprintf(buffer, "\nEvents:\n")

	for _, eventName := range EventNames() {
		fmt.Fprintf(buffer, "  %s (%s)\n", eventName, events[eventName].Name)
	}

	fmt.Fprintln(buffer, "\nNotes:")
	fmt.Fprintln(buffer, "  CasperJS (http://casperjs.org/) is required for authentication.")

	kingpin.UsageTemplate(buffer.String()).Version("2.0").Author("Marcel Jackwerth")
}

func main() {
	configure()
	kingpin.Parse()

	authenticator = NewAuthenticator()
	sessions := SessionList(FindSessions())
	sort.Sort(sessions)

	log.Printf("Found %d sessions.\n", len(sessions))

	for _, session := range sessions {
		if *list {
			fmt.Println(session.String())
			continue
		}

		if !*noVideos {
			DownloadVideo(session)
		}

		if !*noSlides {
			DownloadSlides(session)
		}
	}
}
