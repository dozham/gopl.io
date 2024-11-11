package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	occurrences := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, occurrences)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, occurrences)
			f.Close()
		}
	}
	for line, filesList := range occurrences {
		fmt.Printf("%v\n%s\n", filesList, line)
	}
}

func countLines(f *os.File, occurrences map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		occurrences[input.Text()] = append(occurrences[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
