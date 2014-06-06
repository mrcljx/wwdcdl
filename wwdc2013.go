package main

func init() {
	events["wwdc2013"] = func() []*Session {
		return getSessionsFromUrl("https://developer.apple.com/videos/wwdc/2013/")
	}
}
