package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getDB() *mgo.Database {
	session, err := mgo.Dial("mongodb://127.0.0.1:4001")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("meteor") //root user is created in the admin authentication database and given the role of root.
	return db
}

func main() {
	var mydb = getDB()
	// mydb.Login("root", "123456")
	var m bson.Raw
	c := mydb.C("custom_emoji.files")
	err := c.Find(nil).One(&m)

	check(err)
	fmt.Printf("%#v", m)

	c.Insert(m)

	// for key, value := range m {
	// 	xt := reflect.TypeOf(value).Kind()
	// 	switch xt {
	// 	case reflect.String:
	// 		fmt.Println(key, value)
	// 	case reflect.Slice:
	// 		for k, v := range value {
	// 			fmt.Println(key, value)
	// 		}
	// 	}
	// }
	//mongodb://127.0.0.1:4001/meteor?compressors=disabled&gssapiServiceName=mongodb
}
func check(err error) {
	if err != nil {
		print(err)
	}
}
