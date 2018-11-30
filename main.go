package main

type validate struct {
	i string
	r bool
	e string
}

func validator(u string) validate {
	var n validate
	n.i = u
	n.r = true
	n.e = ""
	return n
}

func main() {
	newInput := validator("golangV$alidation")
	println(newInput.Required().isSpecialCharacter().check())
}
