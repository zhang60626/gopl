package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		PopCount(uint64(8888888888888888))
	}
	fmt.Println(time.Since(start))
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	// TODO
}
