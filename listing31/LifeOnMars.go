package main

import (
	"fmt"
	"sync"
	"image"
	"log"
)

type MarsGrid struct {
	mu sync.Mutex
	size image.Point
}

func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()

}

func NewMarsGrid(size image.Point) *MarsGrid {
	return $MarsGrid{
		size: size,
	}
}

type Occupier struct {
	var cell image.Point
}

func (g *Occupier) Move(p image.Point) bool {

}

