package main

import (
	"fmt"
	"net/http"
	"os"
)

var test = 0

func hello(w http.ResponseWriter, req *http.Request)  {
	_, err := fmt.Fprintf(w, "%d", test)
	test++
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
