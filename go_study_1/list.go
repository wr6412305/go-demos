package main

import (
	"container/list"
	"fmt"
)

func list1() {
	nums := list.New()
	nums.PushBack(1)
	nums.PushBack(2)
	nums.PushBack(3)

	for e := nums.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func printList(coll *list.List) {
	for e := coll.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func printListR(coll *list.List) {
	for e := coll.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func list2() {
	coll := list.New()
	coll.PushBack(1)
	coll.PushBack("Gopher")
	coll.PushFront("Cynhard")
	two := coll.PushFront(2)

	before2 := coll.InsertBefore("Before2", two)
	after2 := coll.InsertAfter("After2", two)

	coll.MoveAfter(before2, two)
	coll.MoveBefore(after2, two)
	coll.MoveToFront(before2)
	coll.MoveToBack(after2)

	coll2 := list.New()
	coll2.PushBack(3)
	coll2.PushBack("Tomcat")

	coll.PushBackList(coll2)
	coll.PushFrontList(coll2)

	printList(coll)
	printListR(coll)

	fmt.Println(coll.Front().Value)
	fmt.Println(coll.Back().Value)
	fmt.Println(coll.Len())

	coll.Remove(two)
	printList(coll)

	coll.Init()
	printList(coll)
}
