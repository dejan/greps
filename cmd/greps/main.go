package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const (
	clear          = "\r                   \r"
	fgCyan         = "\033[0;36m"
	bgRed          = "\033[0;41m"
	reset          = "\033[0m"
	infoColor      = fgCyan
	highlightColor = bgRed
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
			ht := reg.ReplaceAllStringFunc(text, highlight)
			fmt.Print(clear, ht, "\n")
		} else {
			counter++
			fmt.Print(infoColor, "\r", counter, " lines skipped", reset)
		}
	}
	fmt.Println()
}

func highlight(s string) string {
	return highlightColor + s + reset
}
