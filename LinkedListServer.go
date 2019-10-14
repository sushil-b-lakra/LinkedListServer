package main

import "fmt"

type MyNode struct {
	data  int
	pNext *MyNode
}

func main() {

	var nodeptr *MyNode
	//nodeptr := &MyNode{32, nil}

	fmt.Println("Before Adding Node:")
	nodeptr.print()
	addnode(&nodeptr, createnode(36))
	addnode(&nodeptr, createnode(34))
	addnode(&nodeptr, createnode(39))
	addnode(&nodeptr, createnode(40))
	fmt.Println("After Adding Node:")
	nodeptr.print()
	fmt.Println("Sushil Bimal Lakra")
}

func (list *MyNode) print() {
	if list == nil {
		fmt.Println("List is Empty!")
	}
	for list != nil {
		fmt.Printf("%p : { %d , %p}\n", list, list.data, list.pNext)
		list = list.pNext
	}
}

func addnode(list **MyNode, node *MyNode) bool {
	if list == nil {
		return false
	}
	localnode := *list
	if localnode == nil {
		*list = node
	} else {
		var lastnode *MyNode
		for localnode != nil {
			lastnode = localnode
			localnode = localnode.pNext
		}
		lastnode.pNext = node
	}
	return true
}

func createnode(data int) *MyNode {
	node := MyNode{data, nil}
	return &node
}
