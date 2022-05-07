package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world!")
	if err != nil {
		panic("error!")
	}
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("panic!!")
		}
	}()

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic("error!")
	}
}
