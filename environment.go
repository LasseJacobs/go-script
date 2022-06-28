package main

type Environment struct {
	values map[string]Any
}

func NewEnvironment() *Environment {
	return &Environment{values: make(map[string]Any)}
}

func (env *Environment) define(name string, value Any) {
	env.values[name] = value
}

func (env *Environment) get(name Token) Any {
	if v, ok := env.values[name.Lexeme]; ok {
		return v
	}
	panic(NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'."))
}
