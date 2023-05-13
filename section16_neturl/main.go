package main

import (
	"fmt"
	"net/url"
)

func main() {


	url2 := &url.URL{}
	url2.Scheme = "https:"
	url2.Host = "goole.com"
	q2 := url2.Query()
	q2.Set("q", "Golang")
	url2.RawQuery = q2.Encode()

	fmt.Println(url2)
}
