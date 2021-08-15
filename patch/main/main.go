package main

import (
	"context"
	"sync"

	"cj.com/gobot/patch"

	"go.mongodb.org/mongo-driver/bson"
)

const DB = "meteor"

var bgCtx context.Context = context.Background()

// var userclient lun.IClient = &lun.MongoClient{}

func main() {

	userclient := patch.ConnectMongoDB("mongodb://localhost:4001/" + DB)
	defer func() {
		_ = userclient.Disconnect(bgCtx)
	}()
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	db := userclient.Database(DB)
	colls, err := db.ListCollectionNames(bgCtx, bson.D{})
	patch.PanicErr(err)
	for _, colName := range colls {
		wg.Add(1)
		// println("wg add")
		go patch.Replace(userclient, DB, colName, wg, "tsmc.com.tw", "tsmc.com")
		// time.Sleep(1 * time.Second)
	}
	wg.Wait()
}
