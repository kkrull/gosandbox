package main

type Greeter struct{}

func (theGreeter *Greeter) Greet() (string, error) {
	return "hey", nil
}

func NewGreeter() *Greeter {
	return &Greeter{}
}
