package main

import "fmt"

type HelloFunc func(string)

func hello(name string) {
	fmt.Println("Hello, " + name)
}

func say(f HelloFunc) {
	f("Rodson")
}

func greetings(s string) func(string) {
	return func(name string) {
		fmt.Println(s, name)
	}
}

func somador() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	h := hello

	h("Nicolas")

	anon := func() {
		fmt.Println("Anon function")
	}
	anon()

	func() {
		fmt.Println("This is like an IIFE from JS")
	}()

	say(hello)

	g := greetings("Hello")
	g("Nicolas")

	s := somador()
	x := s(5)
	y := s(10)
	fmt.Println(x, y, s(10))
}
