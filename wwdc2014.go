package main

import (
	"strings"
)

func init() {
	events["wwdc2014"] = GetWWDC2014Sessions
}

func GetWWDC2014Sessions() []*Session {
	return getSessionsFromUrl("https://developer.apple.com/videos/wwdc/2014/")
}

func getSessionsFromUrl(url string) []*Session {
	document, _ := ParseHtmlAt(url)
	sessions := make([]*Session, 0)

	for _, sessionNode := range FindNodes(document, "li.session") {
		titleNode := FindNode(sessionNode, "li.title")
		title := InnerHtml(titleNode)
		title = strings.Replace(title, "&#39;", "’", -1)
		sessionNumber := strings.SplitN(GetAttrValue(sessionNode, "id"), "-", 2)[0]
		session := NewSession(sessionNumber, title)

		for _, downloadNode := range FindNodes(sessionNode, ".download a") {
			downloadType := strings.ToLower(InnerHtml(downloadNode))
			downloadUrl := GetAttrValue(downloadNode, "href")
			session.downloads[downloadType] = downloadUrl
		}

		sessions = append(sessions, session)
	}

	return sessions
}
