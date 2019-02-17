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
		name := countLines(os.Stdin, counts)
		if len(name) != 0 {
			fmt.Println("name:", name)
		}
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			name := countLines(f, counts)
			if len(name) != 0 {
				fmt.Println("name:", name)
			}
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) (dupfile string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		t := input.Text()
		counts[t]++
		if counts[t] > 1 && len(dupfile) == 0 {
			dupfile = f.Name()
		}
	}
	return
}
