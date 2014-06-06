package main

import (
	"fmt"
  "flag"
)

var videos bool
var slides bool

func init() {
  flag.BoolVar(&videos, "videos", true, "Download videos")
  flag.BoolVar(&slides, "slides", true, "Download slides/PDFs")
  // TODO: flag for specific sessions
}

func main() {
  flag.Parse()
  
  sessions := FindSessions()
  
  fmt.Printf("Found %d sessions.\n", len(sessions))
  
  for _, session := range sessions {
    if (videos) {
      DownloadVideo(session)
    }
    
    if (slides) {
      DownloadSlides(session)
    }
  }
}