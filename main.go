package main

import "git.taservs.net/ljohnston/example2/ex"

func main() {
	example1 := &Example1{Foo: &ex.T{}}
	example1.Baz()
}
