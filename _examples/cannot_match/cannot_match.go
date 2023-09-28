package main

import (
	"fmt"
	"github.com/nycu-ucr/gonet/http"

	"github.com/h2non/gock"
)

func main() {
	// gock enabled but cannot match any mock
	gock.New("http://httpbin.org").
		Get("/foo").
		Reply(201).
		SetHeader("Server", "gock")

	_, err := http.Get("http://httpbin.org/bar")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
