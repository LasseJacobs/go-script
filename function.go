package main

type Function struct {
	Declaration FunctionStatement
	Closure     *Environment
}

func NewFunction(declaration FunctionStatement, closure *Environment) Function {
	return Function{Declaration: declaration, Closure: closure}
}

func (f Function) Arity() int {
	return len(f.Declaration.Params)
}

func (f Function) Call(interpreter *Interpreter, arguments []Any) Any {
	localEnv := NewEnvironmentWithEnclosing(f.Closure)
	for i, param := range f.Declaration.Params {
		localEnv.define(param.Lexeme, arguments[i])
	}
	return interpreter.executeBlock(f.Declaration.Body, localEnv)
}

func (f Function) String() string {
	return "<fn " + f.Declaration.Name.Lexeme + ">"
}
