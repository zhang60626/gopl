package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		return
	}
	for _, file := range files {
		fd, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(fd, counts)
		fd.Close()
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}

func countLines(fd *os.File, counts map[string]int) {
	input := bufio.NewScanner(fd)
	for input.Scan() {
		counts[input.Text()]++
	}
}
