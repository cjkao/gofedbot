package main

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB = "meteor"

// raw mongo doc raw, pattern string to match
func Ismatch(raw bson.Raw, pattern string) bool {
	return bytes.Index(raw, []byte(pattern)) > 1
}

// replace src string to dest string recursively in value
//
func replaceDeep(raw bson.Raw, src string, dest string) (bson.Raw, error) {
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
	bson.Unmarshal(raw, tmp)
	recDeep(tmp)
	return bson.Marshal(tmp)
}

// replace src to dest in target collection
func replace(collName string, src string, dest string) {
	client := GetMongoDBClient()
	coll := client.Database(DB).Collection(collName)

	cursor, err := coll.Find(bgCtx, bson.D{})
	panicErr(err)
	defer cursor.Close(bgCtx)
	tablePrint := false
	for cursor.Next(bgCtx) {
		raw := cursor.Current
		if Ismatch(raw, src) {
			nRaw, err := replaceDeep(raw, src, dest)
			panicErr(err)
			id := nRaw.Lookup("_id")
			result, err := coll.ReplaceOne(bgCtx, bson.M{"_id": id}, nRaw)
			panicErr(err)
			if !tablePrint {
				tablePrint = true
				println(collName)
			}
			fmt.Printf("%#v\n", result)
		}

	}
}

var bgCtx context.Context = context.Background()
var userclient *mongo.Client

func ConnectMongoDB() {
	bgCtx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	// user Connection database
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:4001/" + DB)
	// Connect to MongoDB
	var err error
	userclient, err = mongo.Connect(bgCtx, clientOptions)
	check(err)
	// Check the connection
	err = userclient.Ping(bgCtx, nil)
	check(err)
	fmt.Println("Connected to meteor MongoDB!")

}

//GetMongoDBClient , return mongo client for CRUD operations
func GetMongoDBClient() *mongo.Client {
	return userclient
}

func main() {
	ConnectMongoDB()
	defer func() {
		_ = userclient.Disconnect(bgCtx)
	}()
	db := userclient.Database(DB)
	colls, err := db.ListCollectionNames(bgCtx, bson.D{})
	panicErr(err)
	for _, colName := range colls {
		replace(colName, "tsmc.com", "tsmc.com.tw")
	}
}

func check(err error) {
	if err != nil {
		fmt.Printf("%#v", err)
	}
}
func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
