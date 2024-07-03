package main

type Greeter struct{}

func (theGreeter *Greeter) Greet() (string, error) {
	return "Hello World!", nil
}

func NewGreeter() *Greeter {
	return &Greeter{}
}
