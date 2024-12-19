package main

import (
	"fmt"
	"os"
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

func runPrompt() {
	fmt.Print("this is runPrompt func\n")
}
