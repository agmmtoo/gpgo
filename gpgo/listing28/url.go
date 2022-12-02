package main

import (
	"fmt"
	"net/url"
	"os"
)

var (
	u = "https://fb .com"
)

func main() {
	parsed, err := url.Parse(u)
	if err != nil {
		fmt.Printf("%#v\n", err)
		if errs, ok := err.(*url.Error); ok {
			fmt.Println("Op: ", errs.Op)
			fmt.Println("URL: ", errs.URL)
			fmt.Println("Err: ", errs.Err)
		}
		os.Exit(1)
	}
	fmt.Println(parsed)
}
