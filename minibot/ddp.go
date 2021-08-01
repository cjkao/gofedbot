package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
	rt "github.com/ruilisi/Rocket.Chat.Go.SDK/realtime"
	// "github.com/stretchr/testify/assert"
)

var (
	client *rt.Client
)

func getLoggedInClient(host string, port string) *rt.Client {

	if client == nil {
		c, _ := rt.NewClient(&url.URL{Host: host + ":" + port}, true)
		// assert.Nil(t, err, "Couldn't create realtime client")
		client = c
	}

	return client
}
func TestClient_SubscribeToMessageStream(c *rt.Client) {

	general := models.Channel{ID: "GENERAL"}
	textToSend := "RealtimeTest"

	message := c.NewMessage(&general, textToSend)
	message.Msg = textToSend

	messageChannel := make(chan models.Message, 1)

	err := c.SubscribeToMessageStream(&general, messageChannel)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		sendAndAssertNoError(c, message)
	}()

	receivedMessage1 := <-messageChannel
	// receivedMessage2 := <-messageChannel
	// receivedMessage3 := <-messageChannel
	fmt.Printf("%#v", receivedMessage1)
	// assertMessage(t, receivedMessage2)
	// assertMessage(t, receivedMessage3)
}

func sendAndAssertNoError(c *rt.Client, message *models.Message) {
	c.SendMessage(message)
}

func TestClient_SubscribeToMessageStream_UnknownChannel(c *rt.Client) {

	channel := models.Channel{ID: "unknown"}
	messageChannel := make(chan models.Message, 1)

	err := c.SubscribeToMessageStream(&channel, messageChannel)

	if err != nil {
		log.Fatal(err)
	}
}
