package main

func init() {
	events["wwdc2012"] = func() []*Session {
		return getSessionsFromUrl("https://developer.apple.com/videos/wwdc/2012/")
	}
}
