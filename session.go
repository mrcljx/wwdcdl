package main

import (
  "fmt"
)

type Session struct {
  Number, Title string
  downloads map[string]string
}

func NewSession(number string, title string) *Session {
  return &Session{
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