// ch1-1.4 prints the name of files which text ppear more than once
// in the input.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	repetition := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			for _, n := range counts {
				if n > 1 {
					repetition[arg]++
				}
			}
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	for line := range repetition {
		fmt.Printf("Name of repetition file: %s\n", line)
	}
}

func countLines(f *os.File, counts map[string]int) {
	//fmt.Printf("countLines Begin!\n")
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//fmt.Printf("countLines End!\n")
}
