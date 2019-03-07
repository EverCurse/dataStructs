/*
单向链表的常用操作
*/
package linkedlist

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}
type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (list *LinkedList) New() {
	//初始化链表
	list.Head = nil
	list.Tail = nil
	list.Size = 0
}

func (list *LinkedList) Append(node *Node) {
	//链表在尾部添加node
	node.next = nil
	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		list.Size += 1
		return
	}
	oldTail := list.Tail
	oldTail.next = node
	list.Tail = node
	list.Size += 1
}

func (list *LinkedList) Print() {
	//遍历链表
	n := list.Head
	res := make([]interface{}, 0)
	for i := 0; i < list.Size; i++ {
		res = append(res, n.value)
		n = n.next
	}
	fmt.Println(res)
}

func (list *LinkedList) Get(index int) (node *Node) {
	//获取链表指定位置的node
	if index >= list.Size {
		return nil
	}
	n := list.Head
	for i := 0; i < index; i++ {
		n = n.next
	}
	return n
}

func (list *LinkedList) Insert(index uint, node *Node) bool {
	//链表插入
	if node == nil || list == nil || list.Size == 0 || int(index) > list.Size {
		return false
	}

	if int(index) == 0 {
		node.next = list.Head
		list.Head = node
		return true
	}

	n := list.Head
	for i := 1; i < int(index); i++ {
		n = n.next
	}
	node.next = n.next
	n.next = node
	list.Size += 1
	return true
}

func (list *LinkedList) Del(index uint) bool {
	//链表删除node
	if list == nil || list.Size == 0 || int(index) > list.Size {
		return false
	}

	if int(index) == 0 {
		list.Head = list.Head.next
		list.Size -= 1
		return true
	}

	n := list.Head
	for i := 1; i < int(index); i++ {
		n = n.next
	}
	n.next = n.next.next
	list.Size -= 1
	return true
}
