package main

import (
	"fmt"
	"net/url"
)

func main() {
	// u, _ := url.Parse("http://example.com/search?a=1&b=2#top")

	// fmt.Println(u.Scheme)
	// fmt.Println(u.Host)
	// fmt.Println(u.Path)
	// fmt.Println(u.RawQuery)
	// fmt.Println(u.Fragment)

	// fmt.Println(u.Query())

	// url := &url.URL{}

	// url.Scheme = "https:"
	// url.Host = "google.com"
	// q := url.Query()
	// q.Set("q", "Golang")

	// url.RawQuery = q.Encode()
	// fmt.Println(url)

	url2 := &url.URL{}
	url2.Scheme = "https:"
	url2.Host = "goole.com"
	q2 := url2.Query()
	q2.Set("q", "Golang")
	url2.RawQuery = q2.Encode()

	fmt.Println(url2)
}
