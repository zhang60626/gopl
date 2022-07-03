package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		PopCount(uint64(8888888888888888))
	}
	fmt.Println(time.Since(start))
}

func PopCount(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int(x & 1)
		x >>= 1
	}
	return count
}
