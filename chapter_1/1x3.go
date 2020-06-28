package main

/*
 * Exc. 1.3:
 *
 * Experiment to measure the difference in running time between our potentially
 * inefficient versions and the one that uses strings.Join.
 */

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	stringsJoin(os.Args)
	rangeLoop(os.Args)
	forLoop(os.Args)
  // Todo: actually measure the difference in running time.
}

func stringsJoin(args []string) {
	fmt.Printf("strings.Join func called:\n%s\n",
		strings.Join(args, " "))
}

func rangeLoop(args []string) {
	fmt.Println("range loop func called:")
	for _, v := range args {
		fmt.Printf("%s ", v)
	}
  fmt.Println("")
}

func forLoop(args []string) {
	fmt.Println("for loop func called:")
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}
