package main

import (
	"net/url"

	"github.com/cjkao/Rocket.Chat.Go.SDK/rest"
)

func InitRcServer(b bool) {
	var _url url.URL
	host := "example.cj:4000"
	_url = url.URL{
		Host: host,
	}
	RestClient = rest.NewClient(&_url, false)
}
func GetSubscribed() {
	RestClient.GetJoinedChannels(nil)
}

var RestClient *rest.Client
