package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string `json:"name" bson:"name"`
	Phone string `json:"phone" bsin:"phone"`
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		fmt.Printf("main err: %v\n", err)
		return
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("runoob").C("runoob")
	err = c.Insert(&Person{"liangjisheng", "48574934745"},
		&Person{"ljs", "376546464"})
	if err != nil {
		fmt.Printf("main insert err: %v\n", err)
		return
	}

	results := []*Person{}
	err = c.Find(bson.M{"name": "ljs"}).All(&results)
	if err != nil {
		fmt.Printf("main find err: %v\n", err)
		return
	}

	if len(results) > 0 {
		for i, result := range results {
			fmt.Printf("data:%d name:%s phone:%s\n", i, result.Name, result.Phone)
		}
	} else {
		fmt.Println("not find data")
	}
}
