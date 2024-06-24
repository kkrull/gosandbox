package ginkgohello

import "fmt"

type Greeter struct {
	Name string
}

func (greeter Greeter) Greet() string {
	var name string
	if greeter.Name != "" {
		name = greeter.Name
	} else {
		name = "World"
	}

	return fmt.Sprintf("Hello %s!", name)
}
