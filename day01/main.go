package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var input io.Reader

	if len(os.Args) > 1 {
		var err error
		file := os.Args[1]
		input, err = os.Open(file)
		if err != nil {
			return err
		}
	} else {
		input = os.Stdin
	}

	data, err := io.ReadAll(input)
	if err != nil {
		return err
	}
	data = bytes.TrimSpace(data)

	lines := strings.Split(string(data), "\n")

	var total int
	for _, line := range lines {
		var (
			first int
			last  int
		)
		foundFirst := false
		for i := range line {
			n, found := toInt(line[i:], false)
			if found && !foundFirst {
				first = n
				foundFirst = true
			}
			if found {
				last = n
			}
		}

		result := first*10 + last
		total += result
	}
	fmt.Println(total)

	var total2 int
	for _, line := range lines {
		var (
			first int
			last  int
		)
		foundFirst := false
		for i := range line {
			n, found := toInt(line[i:], true)
			if found && !foundFirst {
				first = n
				foundFirst = true
			}
			if found {
				last = n
			}
		}

		result := first*10 + last
		total2 += result
	}
	fmt.Println(total2)
	return nil
}

func toInt(input string, includeStr bool) (int, bool) {
	if input == "" {
		return 0, false
	}

	if unicode.IsDigit(rune(input[0])) {
		n, _ := strconv.Atoi(string(input[0]))
		return n, true
	}

	if !includeStr {
		return 0, false
	}

	switch {
	case strings.HasPrefix(input, "one"):
		return 1, true
	case strings.HasPrefix(input, "two"):
		return 2, true
	case strings.HasPrefix(input, "three"):
		return 3, true
	case strings.HasPrefix(input, "four"):
		return 4, true
	case strings.HasPrefix(input, "five"):
		return 5, true
	case strings.HasPrefix(input, "six"):
		return 6, true
	case strings.HasPrefix(input, "seven"):
		return 7, true
	case strings.HasPrefix(input, "eight"):
		return 8, true
	case strings.HasPrefix(input, "nine"):
		return 9, true
	default:
		return 0, false
	}
}
