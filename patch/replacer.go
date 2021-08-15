package patch

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	// "github.com/256dpi/lungo"
	"github.com/256dpi/lungo"
	lun "github.com/256dpi/lungo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var bgCtx context.Context = context.Background()

func Check(err error) {
	if err != nil {
		fmt.Printf("%#v", err)
	}
}
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

// raw mongo doc raw, pattern string to match
// return true, if match any string
// func Ismatch(raw *bson.Raw, pattern string) bool {
// 	return bytes.Index(*raw, []byte(pattern)) > 1
// }

// the input assume there is already exist something that
// replace src string to dest string recursively in value
// return updated bson
func replaceDeep(raw *bson.Raw, src string, dest string) (bson.Raw, error) {
	if bytes.Index(*raw, []byte(src)) < 0 {
		return nil, nil
	}
	var recDeep func(node map[string]interface{})
	recDeep = func(node map[string]interface{}) {
		for k, v := range node {
			switch v.(type) {
			case string: //fix me
				var s string = v.(string)
				if strings.Contains(s, src) {
					node[k] = strings.Replace(s, src, dest, -1)
				}
			case map[string]interface{}:
				nodeNew := v.(map[string]interface{})
				recDeep(nodeNew)
			}
		}
	}
	tmp := make(map[string]interface{})
	err := bson.Unmarshal(*raw, tmp)
	if err != nil {
		panic(err)
	}
	recDeep(tmp)
	return bson.Marshal(tmp)
}

// Replace src to dest in target collection
// return number of record updated
func Replace(client lun.IClient, db, collName string, wg *sync.WaitGroup, src, dest string) int {
	// client := GetMongoDBClient()
	coll := client.Database(db).Collection(collName)
	cursor, err := coll.Find(bgCtx, bson.D{})
	PanicErr(err)
	defer cursor.Close(bgCtx)
	tablePrint := false
	counter := 0
	for cursor.Next(bgCtx) {
		var raw bson.Raw
		cursor.Decode(&raw)
		// if Ismatch(&raw, src) {
		nRaw, err := replaceDeep(&raw, src, dest)
		PanicErr(err)
		if nRaw == nil {
			continue
		}

		id := nRaw.Lookup("_id")
		result, err := coll.ReplaceOne(bgCtx, bson.M{"_id": id}, nRaw)
		counter += int(result.ModifiedCount)
		PanicErr(err)
		if !tablePrint {
			tablePrint = true
			println(collName)
			fmt.Printf("%#v\n", result)
		}

	}
	// println("wg done")
	wg.Done()
	return counter
}

func ConnectMongoDB(dbstr string) lun.IClient {
	bgCtx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	clientOptions := options.Client().ApplyURI(dbstr)
	// Connect to MongoDB
	monClient, err := lungo.Connect(bgCtx, clientOptions)
	Check(err)
	// Check the connection
	err = monClient.Ping(bgCtx, nil)
	Check(err)
	fmt.Println("Connected to meteor MongoDB!")
	return monClient
}

//GetMongoDBClient , return mongo client for CRUD operations
// func GetMongoDBClient() lun.IClient {
// 	return ConnectMongoDB()
// }
