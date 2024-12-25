package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: jlox [script]")
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Failed to read the file %s: %w", path, err)
	}
	content := string(bytes)
	fmt.Println(content)
	return nil
}

func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("Failed to read input : %w", err)
		}
		// convert CRLF to LF

		line = strings.TrimSuffix(line, "\r\n") // Handles CRLF (Windows)
		line = strings.TrimSuffix(line, "\n")   // Handles LF (Unix)
		if len(line) == 0 {
			break
		}
		fmt.Println(line)
		// run(line)
	}
	return nil
}
