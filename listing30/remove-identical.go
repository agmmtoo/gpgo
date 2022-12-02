package main

import (
	"fmt"
	//"strings"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	s := []string{
		"Hi",
		"Hello",
		"Hi",
		"Hello",
		"Good to see you here",
	}

	go sourceGR(c0, s)
	go filterGR(c0, c1)
	printGR(c1)
}

func sourceGR(downstream chan string, ss []string) {
	for _, s := range ss {
		downstream <- s
	}
	close(downstream)
}

func filterGR(upstream, downstream chan string) {
	dict := make(map[string]bool)

	for s := range upstream {
		//s = strings.ToUpperCase(s)
		if !dict[s] {
			dict[s] = true
		}
	}
	for k := range dict {
		downstream <- k
	}
	close(downstream)
}

func printGR(downstream chan string) {
	for s := range downstream {
		fmt.Println(s)
	}
}
