package main

import (
	"fmt"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "TestPage", Body: []byte("Hello world!")}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

