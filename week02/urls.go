package main

import (
	"flag"
	"fmt"
	"net/url"
)

func main() {

	urlPtr := flag.String("url", "", "url to parse")
	flag.Parse()

	u, err := url.Parse(*urlPtr)
	if err != nil {
		fmt.Println("Failed to parse URL ", err)
		return
	}
	fmt.Println("Scheme: ", u.Scheme)
	fmt.Println("Hostname: ", u.Host)
	fmt.Println("Path: ", u.Path)
}
