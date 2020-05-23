package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

//Route describes Traffic
type LinkedListRoute struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

//LinkedList lists handlers for LinkedList Routes
var LinkedListRoutes = []LinkedListRoute{
	LinkedListRoute{"defaulthandler", strings.ToUpper("Get"), "/", defaulthandler},
	LinkedListRoute{"fetchlinkedlistnodeshandler", strings.ToUpper("Get"), "/LinkedListNodes", fetchlinkedlistnodeshandler},
	LinkedListRoute{"addlinkedlistnodehandler", strings.ToUpper("Post"), "/LinkedListNodes", addlinkedlistnodehandler},
	LinkedListRoute{"deletelinkedlistnodehandler", strings.ToUpper("Delete"), "/LinkedListNodes", deletelinkedlistnodehandler},
}

func main() {

	/* This is main function for Linked List Server Program */

        /* Initialize Linked List */
	InitializeLinkedList()

	/* Initializaing Http Server in a separate Go Routine */
	go StartHttpServer()

	/* Initializaing Http Client Command Line Interface */
	StartHttpClientCLI()
}

func StartHttpServer() {

	LinkedListRouter := NewLinkedListRouter()

	httpserver := &http.Server{
		Addr:           ":8095",
		Handler:        LinkedListRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	/* Start HTTP Server for listening for requests */
	err := httpserver.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func NewLinkedListRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range LinkedListRoutes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.Handler)
	}
	return router
}

func defaulthandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<H1>Hello there! This is Linked List HTTP Server!!</H1>")
}

func fetchlinkedlistnodeshandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Linked List contains %d nodes\n", nodeptr.nodecount())
		fmt.Fprintf(w, "All the Nodes in json format are listed as follows: \n")
		jsondata, err := json.Marshal(getallnodesdata(nodeptr))
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%s\n", jsondata)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported Request Method."))
	}
}

func addlinkedlistnodehandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			addnode(&nodeptr, FromJson(body))
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported Request Method."))
	}
}

func deletelinkedlistnodehandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		w.WriteHeader(http.StatusOK)
		deletenode(&nodeptr)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported Request Method."))
	}
}

func (node MyNode) ToJson() []byte {
	jsondata, err := json.Marshal(node.MagicData)
	if err != nil {
		panic(err)
	}
	return jsondata
}

func FromJson(jsondata []byte) *MyNode {
	node := MyNode{}
	magicdata := MagicNumber{}
	err := json.Unmarshal(jsondata, &magicdata)
	if err != nil {
		panic(err)
	}
	node.MagicData.Data = magicdata.Data
	node.MagicData.Description = magicdata.Description
	return &node
}

func (list *MyNode) fprint(w http.ResponseWriter) {
	if list == nil {
		fmt.Fprintf(w, "List is Empty!")
	} else {
		fmt.Fprintf(w, "\nNode Addresss : { Node Data , Node Data Description, Node NextPtr}\n")
		fmt.Fprintf(w, "====================================================================\n")
	}

	for list != nil {
		fmt.Fprintf(w, "%p : { %s , %s , %p}\n", list, list.MagicData.Data, list.MagicData.Description, list.pNext)
		list = list.pNext
	}
}

func (list *MyNode) fprintjson(w http.ResponseWriter) {
	if list == nil {
		fmt.Fprintf(w, "List is Empty!")
	}

	for list != nil {
		fmt.Fprintf(w, "%s\n", list.ToJson())
		list = list.pNext
	}
}
