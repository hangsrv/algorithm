package data_structure

import (
	"container/list"
	"fmt"
)

func InitList() {
	l := list.New()

	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\t", e.Value)
	}

	fmt.Println()

}
