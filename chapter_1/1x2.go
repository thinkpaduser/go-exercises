package main

/*
 * Exc. 1.2:
 *
 * Modify the 'echo' program to print the index and value of each of its
 * arguments, one per line.
 */

import (
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args {
		fmt.Printf("%d: %s\n", k, v)
	}
}
