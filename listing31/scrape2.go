package main

import (
	"fmt"
	"sync"
	"time"
)

type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited) VisitedLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func main() {

	v := Visited{visited: make(map[string]int)}

	go spider1(v)
	go spider2(v)

	time.Sleep(time.Second)
	fmt.Println(v)
}

func spider1(v Visited) {
	v.VisitedLink("https://google.com")
}

func spider2(v Visited) {
	v.VisitedLink("https://youtube.com")
}
