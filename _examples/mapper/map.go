package main

import (
	"fmt"
	"github.com/nycu-ucr/gonet/http"
	
  "github.com/h2non/gock"
)

func main() {
	defer gock.Disable()

	gock.New("http://httpbin.org").
		Get("/").
		Map(func(req *http.Request) *http.Request { req.URL.Host = "httpbin.org"; return req }).
		Map(func(req *http.Request) *http.Request { req.URL.Path = "/"; return req }).
		Reply(204).
	  Map(func(res *http.Response) *http.Response { res.StatusCode = 404; return res }).
		SetHeader("Server", "gock")

	res, err := http.Get("http://httpbin.org/get")
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Server header: %s\n", res.Header.Get("Server"))
}
