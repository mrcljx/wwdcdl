package main

import (
  "strings"
  "log"
  "sync"
)

func init() {
	RegisterEvent(&Event{
		Id: "wwdc2015",
		Name: "WWDC 2015",
		Resolver: func(event *Event) []*Session {
			return getSessionsFromUrl2015(event, "https://developer.apple.com/videos/wwdc/2015/")
		},
	})
}

func getSessionsFromUrl2015(event *Event, url string) []*Session {
	document, _ := ParseHtmlAt(url)
	sessions := make([]*Session, 0)

  log.Println("Getting sessions for WWDC 2015 will take a few seconds because of the website's structure...")

  var waitGroup sync.WaitGroup

	for _, sessionNode := range FindNodes(document, "#all_videos .list_videos li") {
		titleNode := FindNode(sessionNode, "a")
		title := InnerHtml(titleNode)
		title = strings.Replace(title, "&#39;", "’", -1)
		href := GetAttrValue(titleNode, "href")
		sessionNumber := strings.SplitN(href, "=", 2)[1]
		session := NewSession(event, sessionNumber, title)
		sessions = append(sessions, session)

		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			videoDocumentUrl := url + href
			videoDocument, _ := ParseHtmlAt(videoDocumentUrl)

			for _, downloadNode := range FindNodes(videoDocument, "section .text-right a") {
				downloadType := strings.ToLower(InnerHtml(downloadNode))
				downloadUrl := GetAttrValue(downloadNode, "href")
				session.downloads[downloadType] = downloadUrl
			}
		}()
	}

	waitGroup.Wait()

	return sessions
}
