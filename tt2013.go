package main

func init() {
	RegisterEvent(&Event{
		Id: "tt2013",
		Name: "Tech-Talks 2013",
		Resolver: func(event *Event) []*Session {
			return getSessionsFromUrl(event, "https://developer.apple.com/tech-talks/videos/")
		},
	})
}
