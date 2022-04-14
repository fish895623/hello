package main

import "github.com/fish895623/hello/modules"

func main() {
	var a Dog.Animal = Dog.NewDog("asdf")
	println(a.Sounds())
}
