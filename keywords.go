package main

var keywords map[string]TokenType

func init() {
	keywords = make(map[string]TokenType)

	keywords["and"] = TT_AND
	keywords["class"] = TT_CLASS
	keywords["else"] = TT_ELSE
	keywords["false"] = TT_FALSE
	keywords["for"] = TT_FOR
	keywords["fun"] = TT_FUN
	keywords["if"] = TT_IF
	keywords["nil"] = TT_NIL
	keywords["or"] = TT_OR
	keywords["print"] = TT_PRINT
	keywords["return"] = TT_RETURN
	keywords["super"] = TT_SUPER
	keywords["this"] = TT_THIS
	keywords["true"] = TT_TRUE
	keywords["var"] = TT_VAR
	keywords["while"] = TT_WHILE
}
