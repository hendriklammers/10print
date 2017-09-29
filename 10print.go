package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for {
		if random.Float64() <= 0.5 {
			fmt.Print("\u2571")
		} else {
			fmt.Print("\u2572")
		}
		time.Sleep(time.Millisecond * 10)
	}
}
