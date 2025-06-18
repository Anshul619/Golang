package main

// Gorilla Mux is a powerful and flexible HTTP router for Golang, providing advanced features such as route variables, middleware, and optimized performance

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", myHandler)
	log.Fatal(http.ListenAndServe(":8000", r))

	log.Println("never get here")
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	go asynchronousRequest()
	w.Write([]byte("Gorilla!\n"))
}

func asynchronousRequest() {
	amt := time.Duration(1000)
	time.Sleep(time.Millisecond * amt)
	log.Println("Asynchronous request finished")
}
