package main

import (
	"fmt"
	"strings"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	s := "Sometimes it’s easier to operate on words than on sentences."

	go sourceGR(c0, s)
	go filterGR(c0, c1)
	printGR(c1)
}

func sourceGR(downstream chan string, s string) {
		downstream <- s
	close(downstream)
}

func filterGR(upstream, downstream chan string) {
	for s := range upstream {
		for _, w := range strings.Fields(s) {
			downstream <- w
		}
	}
	close(downstream)
}

func printGR(downstream chan string) {
	for s := range downstream {
		fmt.Println(s)
	}
}
