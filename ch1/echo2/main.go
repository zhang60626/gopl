// print command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	for _, arg := range os.Args {
		s += " " + arg
	}
	fmt.Println(s)
}
