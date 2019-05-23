package main

import "fmt"

type user struct {
	ID     int64
	Name   string
	Emails []string
}

func (u user) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.ID, u.Name, u.Emails)
}

type story struct {
	ID       int64
	Title    string
	AuthorID int64
	Author   *user
}

func (s story) String() string {
	return fmt.Sprintf("Story<%d %s %d>", s.ID, s.Title, s.AuthorID)
}
