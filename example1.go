package main

// super important comment

import "fmt"

type Foo interface {
	Bar()
}

type Example1 struct {
	Foo
}

func (e *Example1) Baz() {
	fmt.Println("Baz")
	e.Foo.Bar()
}
