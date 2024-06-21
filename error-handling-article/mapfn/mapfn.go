package main

import "fmt"

func main() {
	mapped := mapFn("anything", func(a any) any {
		return 42
	})

	fmt.Printf("Mapped: %v\n", mapped)
}

type any interface{}
func mapFn(value any, mapper func(any) any) any {
	return mapper(value)
}
