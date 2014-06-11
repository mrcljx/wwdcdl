package main

import (
	"fmt"
	"strings"
	"sort"
)

type Session struct {
	Event *Event
	Number, Title string
	downloads map[string]string
}

func NewSession(event *Event, number string, title string) *Session {
	return &Session{
		event,
		number,
		title,
		make(map[string]string),
	}
}

func (s *Session) fileName(extension string) string {
	return fmt.Sprintf("%s - %s%s", s.Number, s.Title, extension)
}

func (s *Session) Slides() (url string, fileName string, ok bool) {
	url, ok = s.downloads["pdf"]
	return url, s.fileName(".pdf"), ok
}

func (s *Session) Video(preferHd bool) (url string, fileName string, ok bool) {
	hdUrl, hasHd := s.downloads["hd"]
	sdUrl, hasSd := s.downloads["sd"]

	if hasHd && (preferHd || !hasSd) {
		return hdUrl, s.fileName(" (HD).mov"), true
	} else if hasSd {
		return sdUrl, s.fileName(" (SD).mov"), true
	} else {
		return "", "", false
	}
}

func (s *Session) String() string {
	downloadList := make([]string, 0)

	for download, _ := range s.downloads {
		downloadList = append(downloadList, strings.ToUpper(download))
	}

	sort.Strings(downloadList)

	return fmt.Sprintf("%s - %s (%s) \n", s.Number, s.Title, strings.Join(downloadList, ", "))
}
