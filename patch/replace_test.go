package patch

import (
	"context"
	"fmt"
	"sync"
	"testing"

	lun "github.com/256dpi/lungo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type pp struct {
	post post `bson:"post,omitempty"`
}
type post struct {
	Title string   `bson:"title,omitempty"`
	Body  string   `bson:"body,omitempty"`
	Cody  []string `bson:"cody,omitempty"`

	ID primitive.ObjectID `bson:"_id,omitempty"`
}

// func TestDB(t *testing.T){
// 	ConnectMongoDB("meteor")
// }
func TestReplace(t *testing.T) {
	ctx := context.Background()
	// prepare options
	opts := lun.Options{
		Store: lun.NewMemoryStore(),
	}
	client, engine, err := lun.Open(ctx, opts)
	if err != nil {
		panic(err)
	}
	defer engine.Close()
	// get db
	db := client.Database("meteor")

	// get collection
	collA := "col1"
	total := 100
	bar := db.Collection(collA)
	pattern := "tsmc.com.tw"
	for i := 0; i < total; i++ {
		_, err = bar.InsertOne(ctx, &post{
			Title: "Hello World!",
			Body:  fmt.Sprintf("%s %d%s", "Hello", i, pattern),
			Cody: []string{
				fmt.Sprintf("%s%d%s", pattern, i, "str"),
				fmt.Sprintf("%s%d", pattern, i),
			},
		})
		if i%10 == 0 {
			bar.InsertOne(ctx, &post{
				Title: "Hello World!",
			})
			bar.InsertOne(ctx, &pp{
				post: post{
					Title: "Hello World!",
					Body:  fmt.Sprintf("%s %d%s", "Hello", i, pattern),
				},
			})
		}

	}
	mm := bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}
	mm["foo"] = bson.M{"foo": "bar", "hello": "world", "pi": "a" + pattern}
	bar.InsertOne(ctx, mm)
	wg := &sync.WaitGroup{}
	// insert post

	if err != nil {
		panic(err)
	}
	type args struct {
		client   lun.IClient
		db       string
		collName string
		wg       *sync.WaitGroup
		src      string
		dest     string
		dry      bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"mem", args{client, "meteor", collA, wg, "tsmc.com.tw", "tsmc.com", true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg.Add(1)
			replaced := Replace(tt.args.client, tt.args.db, tt.args.collName, tt.args.wg, tt.args.src, tt.args.dest, tt.args.dry)
			assert.Equal(t, total+1, replaced)

		})
	}
}

func TestExample(t *testing.T) {
	type post struct {
		Title string `bson:"title"`
	}

	// prepare options
	opts := lun.Options{
		Store: lun.NewMemoryStore(),
	}

	// open database
	client, engine, err := lun.Open(nil, opts)
	if err != nil {
		panic(err)
	}

	// ensure engine is closed
	defer engine.Close()

	// get db
	foo := client.Database("foo")

	// get collection
	bar := foo.Collection("bar")

	// insert post
	_, err = bar.InsertOne(nil, &post{
		Title: "Hello World!",
	})
	if err != nil {
		panic(err)
	}

	// query posts
	csr, err := bar.Find(nil, bson.M{})
	if err != nil {
		panic(err)
	}

	// decode posts
	var posts []post
	err = csr.All(nil, &posts)
	if err != nil {
		panic(err)
	}

	// print documents
	fmt.Printf("%+v", posts)

	// Output:
	// [{Title:Hello World!}]
}
