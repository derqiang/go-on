package ch7

import (
	"fmt"
	"log"
	"net/http"
)

type ServerRunner int64

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

/// database define database structure
type database map[string]dollars

func (db *database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range *db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := (*db)[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no sum item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func (ServerRunner) Run() {

	var db = database{
		"shoes": 50,
		"socks": 5,
	}

	log.Fatal(http.ListenAndServe("localhost:8080", &db))
}
