package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/rest"
)

func main() {
	InitRcServer(true)
	login := rest.AuthLoginRequest{
		User:     "b",
		Password: "b",
	}
	loginres, err := RestClient.AuthLogin(login)
	if err != nil {
		print(err)
		panic(err)
	}

	PrettyPrtint(loginres.Data)
	parms := make(map[string][]string)
	parms["roomName"] = []string{"general"}
	parms["unreads"] = []string{"true"}
	parms["latest"] = []string{"2021-07-30"}
	parms["oldest"] = []string{"2021-07-01"}
	parms["inclusive"] = []string{"true"}
	resp, err := RestClient.ChannelHistory(parms)
	if err != nil {
		log.Fatal(err)
	}
	PrettyPrtint(resp)
}
func PrettyPrtint(x interface{}) {
	empJSON, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("x==>\n %s\n", string(empJSON))

}
