package main

func init() {
	RegisterEvent(&Event{
		Id: "wwdc2013",
		Name: "WWDC 2013",
		Resolver: func() []*Session {
			return getSessionsFromUrl("https://developer.apple.com/videos/wwdc/2013/")
		},
	})
}