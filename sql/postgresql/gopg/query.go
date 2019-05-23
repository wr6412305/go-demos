package main

import (
	"fmt"

	"github.com/go-pg/pg"
)

func createData() {
	db := connect()
	defer db.Close()

	u := &user{
		Name:   "admin",
		Emails: []string{"admin1@admin", "admin2@admin"},
	}
	_, err := db.QueryOne(&user{}, "insert into users (name, emails) values (?name, ?emails) returning id", u)
	if err != nil {
		panic(err)
	}

	story1 := &story{
		Title:    "Cool story",
		AuthorID: u.ID,
	}
	_, err = db.QueryOne(&story{}, "insert into stories (title, author_id) values (?title, ?author_id) returning id", story1)
	if err != nil {
		panic(err)
	}

	u = &user{
		Name:   "root",
		Emails: []string{"root1@root", "root2@root"},
	}
	_, err = db.QueryOne(&user{}, "insert into users (name, emails) values (?name, ?emails) returning id", u)
	if err != nil {
		panic(err)
	}
}

func query() {
	db := connect()
	defer db.Close()

	var users []user
	_, err := db.Query(&users, "select * from users")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", users)

	var u user
	_, err = db.QueryOne(&u, "select * from users where id = ?", 6)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	var users1 []user
	_, err = db.Query(&users1, "select * from users where id in (?)", pg.In([]int64{6, 7}))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", users1)
}
