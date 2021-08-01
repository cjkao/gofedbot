package main

import (
	"net/url"

	t "cj.com/gobot/testing"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/realtime"
	re "github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
	// "github.com/pyinx/gorocket/common_testing"
	// "github.com/stretchr/testify/assert"
)

var (
	testUserName  string
	testUserEmail string
	testPassword  = "test"
	rocketClient  *re.Client
)

func getDefaultClient() *re.Client {

	if rocketClient == nil {
		testUserEmail = "bottt@b.com"
		testUserName = "botttt"
		rocketClient = getAuthenticatedClient(testUserName, testUserEmail, testPassword)
	}

	return rocketClient
}

func getAuthenticatedClient(name, email, password string) *re.Client {
	client := re.Client{Protocol: t.Protocol, Host: t.Host, Port: t.Port}
	credentials := &models.UserCredentials{Name: name, Email: email, Password: password}

	rtClient, err := realtime.NewClient(&url.URL{Host: t.Host + ":" + t.Port}, true)
	if err != nil {
		panic(err)
	}
	_, regErr := rtClient.RegisterUser(credentials)
	if regErr != nil {
		print(regErr)
	}
	loginErr := client.Login(credentials)
	if loginErr != nil {
		panic(err)
	}
	return &client
}

func findMessage(messages []models.Message, user string, msg string) *models.Message {
	var m *models.Message
	for i := range messages {
		m = &messages[i]
		if m.User.UserName == user && m.Msg == msg {
			return m
		}
	}

	return nil
}

func getChannel(channels []models.Channel, name string) *models.Channel {
	for _, r := range channels {
		if r.Name == name {
			return &r
		}
	}

	return nil
}

func main() {
	clie := getDefaultClient()
	// room, err := clie.CreateDirectMessage("example")
	// print(err)
	ret, err := clie.GetJoinedChannels(nil)
	// clie.
	print(err)
	print(ret)
	// print(room)
}
