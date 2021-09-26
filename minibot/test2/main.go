package main

import (
	"encoding/json"
	"fmt"
	"os"

	rt "github.com/cjkao/Rocket.Chat.Go.SDK/realtime"
)

func main() {
	dat, err := os.ReadFile("aa.json")
	check(err)
	rawJ := make(map[string]interface{})
	json.Unmarshal(dat, rawJ)
	m := rt.GetMessageGabs(rawJ["args"])
	fmt.Printf("%#v", m)
	mod := new(AutoGenerated)
	err = json.Unmarshal(dat, mod)
	check(err)
	print(mod.EventName)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Fields struct {
	EventName string    `json:"eventName"`
	Args      []Message `json:"args"`
}

type Message struct {
	ID        string `json:"_id,omitempty"`
	UpdatedAt struct {
		Date int64 `json:"$date"`
	} `json:"_updatedAt,omitempty"`
	Channels []interface{} `json:"channels,omitempty"`
	Md       []struct {
		Type  string `json:"type"`
		Value []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"value"`
	} `json:"md,omitempty"`
	Mentions []interface{} `json:"mentions,omitempty"`
	Msg      string        `json:"msg,omitempty"`
	Rid      string        `json:"rid,omitempty"`
	Ts       struct {
		Date int64 `json:"$date"`
	} `json:"ts,omitempty"`
	U struct {
		ID       string `json:"_id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"u,omitempty"`
	Urls            []interface{} `json:"urls,omitempty"`
	RoomName        string        `json:"roomName,omitempty"`
	RoomParticipant bool          `json:"roomParticipant,omitempty"`
	RoomType        string        `json:"roomType,omitempty"`
}
type AutoGenerated struct {
	Args      []Message `json:"args"`
	EventName string    `json:"eventName"`
}