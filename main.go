package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong!"))
	fmt.Println("ponged a ping")
}

func main() {
	http.HandleFunc("/ping", handler)

	fmt.Println("ping listening on 0.0.0.0, port 9999")
	err := http.ListenAndServe(":9999", nil)

	if err != nil {
		fmt.Println("Error starting ping server: ", err)
	}
}
