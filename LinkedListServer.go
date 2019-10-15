package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type MyNode struct {
	data  int
	pNext *MyNode
}

var nodeptr *MyNode

func main() {

	//fmt.Println("Before Adding Node:")
	//nodeptr.print()
	fmt.Println("Server is running ... \n")

	addnode(&nodeptr, createnode(36))
	addnode(&nodeptr, createnode(34))
	addnode(&nodeptr, createnode(39))
	addnode(&nodeptr, createnode(40))

	fmt.Println("Initial Linked List is as follows: \n")
	nodeptr.print()

	/* Register Handler functions for APIs */
	http.HandleFunc("/", rootnodehandler)
	http.HandleFunc("/api", rootnodehandler)
	http.HandleFunc("/api/fetchlinkedlistnodes", fetchlinkedlistnodeshandler)
	http.HandleFunc("/api/addlinkedlistnode", addlinkedlistnode)
	http.HandleFunc("/api/deletelinkedlistnode", deletelinkedlistnode)

	/* Start HTTP Server for listening for requests */
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		panic(err)
	}

}

func rootnodehandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<H1>Hello there! This is Cloud Native Server!!</H1>")
}

func fetchlinkedlistnodeshandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Linked List contains %d nodes\n", nodeptr.nodecount())
	fmt.Fprintf(w, "Nodes are listed as follows: \n")
	nodeptr.fprint(w)
}

func addlinkedlistnode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data := r.URL.Query()["data"][0]
	num, _ := strconv.Atoi(data)
	addnode(&nodeptr, createnode(num))
}

func deletelinkedlistnode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	deletenode(&nodeptr)
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

func (list *MyNode) print() {
	if list == nil {
		fmt.Println("List is Empty!")
	} else {
		fmt.Println("Node Addresss : { Node Data , Node NextPtr}")
		fmt.Println("===========================================")
	}
	for list != nil {
		fmt.Printf("%p : { %d , %p}\n", list, list.data, list.pNext)
		list = list.pNext
	}
}

func (list *MyNode) fprint(w http.ResponseWriter) {
	if list == nil {
		fmt.Fprintf(w, "List is Empty!")
	} else {
		fmt.Fprintf(w, "\nNode Addresss : { Node Data , Node NextPtr}\n")
		fmt.Fprintf(w, "===========================================\n")
	}

	for list != nil {
		fmt.Fprintf(w, "%p : { %d , %p}\n", list, list.data, list.pNext)
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

func createnode(data int) *MyNode {
	node := MyNode{data, nil}
	return &node
}
