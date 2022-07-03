package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
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
		counts[file] = make(map[string]int)
		countLines(fd, counts[file])
		fd.Close()
	}
	for file, subcounts := range counts {
		for line, n := range subcounts {
			if n > 1 {
				fmt.Printf("%s\t%s\t%d\n", file, line, n)
			}
		}
	}
}

func countLines(fd *os.File, counts map[string]int) {
	input := bufio.NewScanner(fd)
	for input.Scan() {
		counts[input.Text()]++
	}
}
