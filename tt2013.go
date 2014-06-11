package main

func init() {
	RegisterEvent(&Event{
		Id: "tt2013",
		Name: "Tech-Talks 2013",
		Resolver: func() []*Session {
			return getSessionsFromUrl("https://developer.apple.com/tech-talks/videos/")
		},
	})
}
