package main

import (
	"container/ring"
	"fmt"
)

func ring1() {
	const rLen = 3

	// 创建新的Ring(环形链表)
	r := ring.New(rLen)

	for i := 0; i < rLen; i++ {
		r.Value = i
		r = r.Next()
	}

	fmt.Printf("Length of ring: %d\n", r.Len())

	printRing := func(v interface{}) {
		fmt.Print(v, " ")
	}

	// 打印元素
	r.Do(printRing)
	fmt.Println()

	// 将r之后的第二个元素值乘以2
	r.Move(2).Value = r.Move(2).Value.(int) * 2

	r.Do(printRing)
	fmt.Println()

	// 删除r与r+2之间的元素，即删除r+1
	// 返回删除的元素组成的Ring的指针
	result := r.Link(r.Move(2))

	r.Do(printRing)
	fmt.Println()

	result.Do(printRing)
	fmt.Println()

	another := ring.New(rLen)
	another.Value = 7
	another.Next().Value = 8
	another.Prev().Value = 9

	another.Do(printRing)
	fmt.Println()

	// 插入another到r后面，返回插入前r的下一个与元素
	result = r.Link(another)

	r.Do(printRing)
	fmt.Println()

	result.Do(printRing)
	fmt.Println()

	// 删除r之后的三个元素，返回被删除元素组成的Ring指针
	result = r.Unlink(3)

	r.Do(printRing)
	fmt.Println()

	result.Do(printRing)
	fmt.Println()
}
