//!+

// Dup prints the count and text of lines that appear more than once
// in the input. It also prints the names of the files, in which the
// duplication occurs. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			var keys []string

			for key := range filenames[line] {
				keys = append(keys, key)
			}
			fmt.Printf("\t%s:\n", strings.Join(keys, " "))
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	filename := f.Name()
	for input.Scan() {
		text := input.Text()
		counts[text]++

		if _, ok := filenames[text]; !ok {
			filenames[text] = make(map[string]bool)
		}

		filenames[text][filename] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
