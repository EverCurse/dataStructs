package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := LinkedList{}
	list.New()
	list.Append(&Node{
		value:1,
		next:nil,
	})

	list.Append(&Node{
		value:2,
		next:nil,
	})


	list.Append(&Node{
		value:3,
		next:nil,
	})


	list.Print()

	t.Log(list.Get(2).value)

}


func TestLinkedList_Insert(t *testing.T) {
	list := LinkedList{}
	list.New()
	list.Append(&Node{
		value:10,
		next:nil,
	})

	list.Append(&Node{
		value:20,
		next:nil,
	})


	list.Append(&Node{
		value:30,
		next:nil,
	})

	list.Insert(2,&Node{value:100,next:nil})
	list.Print()
}


func TestLinkedList_Del(t *testing.T) {
	list := LinkedList{}
	list.New()
	list.Append(&Node{
		value:10,
		next:nil,
	})

	list.Append(&Node{
		value:20,
		next:nil,
	})


	list.Append(&Node{
		value:30,
		next:nil,
	})

	list.Del(0)
	list.Print()
}