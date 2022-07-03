package main

type Callable interface {
	Arity() int
	Call(interpreter *Interpreter, arguments []Any) Any
}
