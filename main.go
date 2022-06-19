package main

import (
	"bufio"
	"fmt"
	"os"
)

var hasError bool

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fault(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[line %d] %s: %s", line, where, message)
	hasError = true
}

func run(source string) {
	var scanner = NewScanner(source)
	var tokens = scanner.ScanTokens()
	for _, t := range tokens {
		fmt.Println(t.String())
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		check(err)
		run(line)
		hasError = false
	}
}

func runScript(filename string) {
	bytes, err := os.ReadFile(filename)
	check(err)
	run(string(bytes))
	if hasError == true {
		os.Exit(64)
	}
}

func main2() {
	if len(os.Args) == 1 {
		runPrompt()
	} else if len(os.Args) == 2 {
		runScript(os.Args[1])
	} else {
		println("Usage: lox <script>")
		os.Exit(64)
	}
}
