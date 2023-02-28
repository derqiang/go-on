package ch7

import (
	"log"
	"net/http"
)

type Server2Runner int64

func (sr Server2Runner) Run() {
	parallelRun()
	return
	log.Println("单一Web服务器")
	var db = database{
		"shoes": 50,
		"socks": 5,
	}

	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc((&db).list))
	mux.Handle("/price", http.HandlerFunc((&db).price))
	http.ListenAndServe("localhost:8080", mux)
}

/// 并行
func parallelRun() {
	log.Println("并行默认Web服务器")
	var db = database{
		"shoes": 50,
		"socks": 5,
	}
	http.HandleFunc("/list", (&db).list)
	http.HandleFunc("/price", (&db).price)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	db.ServeHTTP(w, req)
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	db.ServeHTTP(w, req)
}
