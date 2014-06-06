package main

func init() {
	events["tt2013"] = func() []*Session {
		return getSessionsFromUrl("https://developer.apple.com/tech-talks/videos/")
	}
}
