package main

type Environment struct {
	enclosing *Environment
	values    map[string]Any
}

func NewEnvironment() *Environment {
	return &Environment{enclosing: nil, values: make(map[string]Any)}
}

func NewEnvironmentWithEnclosing(enclosing *Environment) *Environment {
	return &Environment{enclosing: enclosing, values: make(map[string]Any)}
}

func (env *Environment) define(name string, value Any) {
	env.values[name] = value
}

func (env *Environment) get(name Token) Any {
	if v, ok := env.values[name.Lexeme]; ok {
		return v
	}
	if env.enclosing != nil {
		return env.enclosing.get(name)
	}
	panic(NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'."))
}

func (env *Environment) getAt(distance int, name string) Any {
	return env.ancestor(distance).values[name]
}

func (env *Environment) ancestor(distance int) *Environment {
	var it = env
	for i := 0; i < distance; i++ {
		it = it.enclosing
	}
	return it
}

func (env *Environment) assign(name Token, value Any) {
	if _, ok := env.values[name.Lexeme]; ok {
		env.values[name.Lexeme] = value
		return
	}
	if env.enclosing != nil {
		env.enclosing.assign(name, value)
		return
	}
	panic(NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'."))
}

func (env *Environment) assignAt(distance int, name Token, value Any) {
	env.ancestor(distance).values[name.Lexeme] = value
}
