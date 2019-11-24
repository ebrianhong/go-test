package main

import (
	"fmt"
	"net/http"
)
var print = fmt.Println

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/headers", Headers)
	http.ListenAndServe(":8080", nil)
}

func HelloServer (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func Headers (w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
