package main

import (
	"github.com/fish895623/hello"
)

func main() {
	var a hello.Animal = hello.NewDog("asdf")
	println(a.Sounds())
}
