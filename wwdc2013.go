package main

func init() {
	RegisterEvent(&Event{
		Id: "wwdc2013",
		Name: "WWDC 2013",
		Resolver: func(event *Event) []*Session {
			return getSessionsFromUrl(event, "https://developer.apple.com/videos/wwdc/2013/")
		},
	})
}