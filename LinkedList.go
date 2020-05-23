package main

import "fmt"

type MagicNumber struct {
	Data        string
	Description string
}

type MyNode struct {
	MagicData MagicNumber
	pNext     *MyNode
}

var nodeptr *MyNode

func InitializeLinkedList() {

	fmt.Printf("Function for Initializaing a sample Linked List\n")
        addnode(&nodeptr, createnode("7", "Days of Week"))
	addnode(&nodeptr, createnode("64", "Boxes in Checker"))
	addnode(&nodeptr, createnode("12", "Months of Year"))
	addnode(&nodeptr, createnode("60", "Seconds in a Minute"))

	fmt.Printf("Initializaing a linked list of size: %d\n", nodeptr.nodecount())
	fmt.Println("Linked List Data is as follows: \n")
	nodeptr.print()
}

func (list *MyNode) nodecount() int {
	len := 0
	if list == nil {
		return 0
	}
	for list != nil {
		list = list.pNext
		len++
	}
	return len
}

func createnode(data string, description string) *MyNode {
	num := MagicNumber{data, description}
	node := MyNode{num, nil}
	return &node
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

func deletenode(list **MyNode) bool {
	if list == nil {
		return false
	}
	lastnode := *list
	if lastnode != nil {
		prevtolastnode := *list
		for lastnode.pNext != nil {
			prevtolastnode = lastnode
			lastnode = lastnode.pNext
		}
		if lastnode == prevtolastnode {
			*list = nil
		} else {
			prevtolastnode.pNext = nil
		}
	}
	return true
}

func (list *MyNode) print() {
	if list == nil {
		fmt.Println("List is Empty!")
	} else {
		fmt.Println("Node Addresss : { Node Data , Node Data Description, Node NextPtr}")
		fmt.Println("==================================================================")
	}
	for list != nil {
		fmt.Printf("%p : { %s , %s , %p}\n", list, list.MagicData.Data, list.MagicData.Description, list.pNext)
		list = list.pNext
	}
}

func getallnodesdata(list *MyNode) []MagicNumber {
	var magicnumdata []MagicNumber
	if list == nil {
		fmt.Println("List is Empty!")
		return magicnumdata
	}
	for list != nil {
		fmt.Printf("%p : { %s , %s , %p}\n", list, list.MagicData.Data, list.MagicData.Description, list.pNext)
		magicnumdata = append(magicnumdata, MagicNumber{list.MagicData.Data, list.MagicData.Description})
		list = list.pNext
	}
	return magicnumdata
}
