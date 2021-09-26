package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/Jeffail/gabs/v2"
	"github.com/cjkao/Rocket.Chat.Go.SDK/models"
	rt "github.com/cjkao/Rocket.Chat.Go.SDK/realtime"
	// "github.com/stretchr/testify/assert"
)

var (
	// rtClient *rt.Client
	user *models.User
)

func main() {
	rtCli := getLoggedInClient("example.cj", "4002")
	PrettyPrint(rtCli)
	SubscribeToMessageStream(rtCli)
}

func getLoggedInClient(host string, port string) (rtClient *rt.Client) {
	c, _ := rt.NewClient(&url.URL{Host: host + ":" + port}, true)
	rtClient = c
	u, error := rtClient.Login(&models.UserCredentials{Username: "b", Password: "b", Email: "b@b.b"})
	check(error)
	user = u
	PrettyPrint(user)
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func SubscribeToMessageStream(c *rt.Client) {
	messageChannelRaw := make(chan map[string]interface{}, 1)
	go func(msgChan chan map[string]interface{}) {
		for i := 0; i < 100; i++ {
			t := <-msgChan
			gc := gabs.Wrap(t)
			en := gc.Path("eventName").Data().(string)
			print(en)
			for k, v := range gc.Path("args.0").ChildrenMap() {
				v.ArrayCount()
				fmt.Printf("%v %v|", k, v)
			}
			for k, v := range gc.Path("args.1").ChildrenMap() {
				v.ArrayCount()
				fmt.Printf("==> %v %v|", k, v)
			}
		}
	}(messageChannelRaw)
	err := c.SubscribeToMyMessagesRaw(messageChannelRaw)
	if err != nil {
		log.Fatal(err)
	}
	done2 := make(chan interface{})
	<-done2

}

// func save(bstr []byte) {
// 	err := os.WriteFile("aa.json", bstr, 0644)
// 	check(err)
// }

// PrettyPrint convert structure as JSON to pretty print string
func PrettyPrint(i interface{}) {
	switch i.(type) {
	case models.Message:
		// t := i.(models.Message)

		s, _ := json.MarshalIndent(i, "\t", "\t")
		fmt.Printf("fffffffformat : %#v\n", string(s))
	default:

		fmt.Printf("nn format : %#v\n", i)

	}
}
