package main

import (
	"fmt"
	"time"
)

func main() {
	future := time.Unix(12622780800, 0)
	fmt.Printf("future %T %[1]v\n", future)
}
