package main

import "time"

type clockFn struct {
}

func (fn clockFn) Arity() int {
	return 0
}

func (fn clockFn) Call(interpreter *Interpreter, arguments []Any) Any {
	return float64(time.Now().UnixMilli())
}

func (fn clockFn) String() string {
	return "<native fn>"
}
