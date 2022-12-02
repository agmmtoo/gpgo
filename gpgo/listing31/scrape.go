package main

import (
	"fmt"
	"math/rand"
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
	c := make(chan Visited)

	go spider1(c)
	go spider2(c)

	fmt.Println(<-c)
}

func spider1(c chan Visited) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	v := <-c
	v.VisitedLink("https://google.com")
}

func spider2(c chan Visited) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	v := <-c
	v.VisitedLink("https://youtube.com")
}
