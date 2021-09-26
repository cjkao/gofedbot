package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/cjkao/Rocket.Chat.Go.SDK/models"
	rest "github.com/cjkao/Rocket.Chat.Go.SDK/rest"
)

func getAuthenticatedClient(name, email, password string) *rest.Client {
	client := rest.Client{Protocol: "HTTP", Host: "example.cj", Port: "4000"}
	credentials := &models.UserCredentials{Name: name, Email: email, Password: password}
	credentials.ID = "DPt636Ki5GpkLpP4f"
	credentials.Token = "TqoqS55rdsL2Y2aCoTi-q7c15Qrew_5_fuoQH-e-562"
	_, err := client.Login(credentials)
	if err != nil {
		panic(err)
	}

	return &client
}

//
//
//Token: TqoqS55rdsL2Y2aCoTi-q7c15Qrew_5_fuoQH-e-562
//Your user Id: DPt636Ki5GpkLpP4f

func main() {
	clie := getAuthenticatedClient("", "", "")
	room, err := clie.CreateDirectMessage("b")
	if err != nil {
		log.Fatal(err)
	}

	// clie.GetDirectory()
	fmt.Printf("room: %#v", room)
	// id := getUserId(clie, "playb@play.cj")
	// fmt.Printf("playb: %s", id)
	// clie.TeamsCreate(&models.TeamsCreateRequest{
	// 	Name:    "X1-XX",
	// 	Type:    1,
	// 	Members: []string{"a", "playx@play.cj."},
	// })

	grpResp, err := clie.CreateGroup(&models.CreateGroupRequest{
		Name:     "Acc1",
		Members:  []string{"b"},
		ReadOnly: false,
	})
	if err == nil {
		// clie.CreateDiscussion(fmt.Sprintf("A-%d", i), &models.Room{ID: grpResp.Group.ID})
		respDis, _ := clie.CreateDiscussion2(&rest.CreateDiscussionInput{Prid: grpResp.Group.ID, TName: fmt.Sprintf("A%d", i), Reply: "AAAAAAAAAAA"})
		fmt.Printf("--->%v", respDis)
	}

	// print(err)
	ret, err := clie.GetJoinedChannels(nil)
	// clie.
	print(err)
	print(ret)
	// print(room)
}
func getRemoteUser(cli *rest.Client, username string) {
	v := url.Values{}
	v["query"] = []string{url.QueryEscape(fmt.Sprintf(`{"type":"users","text":"$s","workspace":"external"}`, username))}
	dir, err := cli.GetDirectory(v)
	if err != nil {
		log.Fatal(err)
	}
	print(dir)
	// dir.Result
}
func getUserId(cli *rest.Client, username string) string {
	v := url.Values{}
	v["username"] = []string{username}
	resp, err := cli.UserInfo(v)
	if err != nil {
		log.Fatal(err)
	}
	return resp.UserInfo.ID
}
