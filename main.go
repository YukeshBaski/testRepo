package main

import "fmt"

func main() {
	fmt.Println("first staging area.")
	genericFunc("hello")
	genericFunc(123)
}

func genericFunc[k string | int](thing k) {
	fmt.Println(thing)
}
