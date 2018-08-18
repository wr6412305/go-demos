package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func test_json1() {
	const jsonStream = `
		{ "Name" : "Ed" , "Text" : "Knock knock." }
		{ "Name" : "Sam" , "Text" : "Who's there?" }
		{ "Name" : "Ed" , "Text" : "Go fmt." }
		{ "Name" : "Sam" , "Text" : "Go fmt who?" }
		{ "Name" : "Ed" , "Text" : "Go fmt yourself!" }
	`
	type Message struct {
		Name, Text string
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

func json_marshal() {
	// Marshal序列化json格式
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)
}

func json_Unmarshal() {
	var jsonBlob = []byte(`[
		{ "Name" : "Platypus" , "Order" : "Monotremata" } , 
        { "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" }
	]`)

	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

func main() {
	test_json1()
	fmt.Println()
	json_marshal()
	fmt.Println()
	json_Unmarshal()

}
