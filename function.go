package main

type Function struct {
	Declaration FunctionStatement
}

func NewFunction(Declaration FunctionStatement) Function {
	return Function{Declaration: Declaration}
}

func (f Function) Arity() int {
	return len(f.Declaration.Params)
}

func (f Function) Call(interpreter *Interpreter, arguments []Any) Any {
	localEnv := NewEnvironmentWithEnclosing(interpreter.globals)
	for i, param := range f.Declaration.Params {
		localEnv.define(param.Lexeme, arguments[i])
	}
	interpreter.executeBlock(f.Declaration.Body, localEnv)
	return nil
}

func (f Function) String() string {
	return "<fn " + f.Declaration.Name.Lexeme + ">"
}
