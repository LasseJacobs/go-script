package main

import (
	"bufio"
	"fmt"
	"os"
)

var hadError bool
var hadRuntimeError bool

var interpreter *Interpreter

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fault(line int, message string) {
	report(line, "", message)
}

func parseFault(token Token, message string) {
	if token.TokenType == TT_EOF {
		report(token.Line, " at end", message)
	} else {
		report(token.Line, " at '"+token.Lexeme+"'", message)
	}
}

func report(line int, where string, message string) {
	fmt.Printf("[line %d] %s: %s\n", line, where, message)
	hadError = true
}

func runtimeFault(err RuntimeError) {
	fmt.Printf("[line %d] %s\n", err.token.Line, err.message)
	hadRuntimeError = true
}

func run(source string) {
	var scanner = NewScanner(source)
	var tokens = scanner.ScanTokens()

	var parser = NewParser(tokens)
	stmts := parser.Parse()

	if hadError {
		return
	}

	resolver := NewResolver(interpreter)
	resolver.Resolve(stmts)

	if hadError {
		return
	}

	interpreter.Interpret(stmts)
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		check(err)
		run(line)
		hadError = false
	}
}

func runScript(filename string) {
	bytes, err := os.ReadFile(filename)
	check(err)
	run(string(bytes))
	if hadError == true {
		os.Exit(64)
	}
	if hadRuntimeError == true {
		os.Exit(70)
	}
}

func main() {
	interpreter = NewInterpreter()
	if len(os.Args) == 1 {
		runPrompt()
	} else if len(os.Args) == 2 {
		runScript(os.Args[1])
	} else {
		println("Usage: lox <script>")
		os.Exit(64)
	}
}
