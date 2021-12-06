package main

import (
	"errors"
	"fmt"
)

func fizz() error {
	return errors.New("FUCK")
}

func main() {
	var f string = "foo"
	fmt.Printf("f is ", f)
	var bar = func(foo string) error {
		fmt.Printf("bar says: %s", foo)
		return nil
	}
	bar("sup")
	fizz()
}

type Car struct {
	Make  string
	Model string
}

func New() *Car {
	return &Car{}
}
