package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const (
	clear   = "\r                   \r"
	skipped = " lines skipped"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Argument missing\n")
		os.Exit(0)
	}
	regexpIn := os.Args[1]
	reg, err := regexp.Compile(regexpIn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid regexp\n")
		os.Exit(0)
	}
	scanner := bufio.NewScanner(os.Stdin)
	var counter int64
	for scanner.Scan() {
		text := scanner.Text()
		if reg.MatchString(text) {
			fmt.Print(clear, text, "\n")
		} else {
			counter++
			fmt.Print("\r", counter, skipped)
		}
	}
	fmt.Println()
}
