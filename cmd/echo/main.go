package main

import (
	isupov "ituniver/third"
	"net/http"
)

func main() {
	http.HandleFunc("/", isupov.Echo)

	http.ListenAndServe(":8880", nil)
}
