package main

import (
	"fmt"
	"github.com/nycu-ucr/gonet/http"

	"github.com/h2non/gock"
)

func main() {
	defer gock.Off()

	gock.New("http://httpbin.org").
		Get("/").
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) { return req.URL.Scheme == "http", nil }).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) { return req.Method == ereq.Method, nil }).
		Reply(204).
		SetHeader("Server", "gock")

	res, err := http.Get("http://httpbin.org/get")
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Server header: %s\n", res.Header.Get("Server"))
}
