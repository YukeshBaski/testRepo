package main

import (
	"fmt"

	"github.com/YukeshBaski/testRepo.git/helper"
)

func main() {
	fmt.Println("first staging area.")
	genericFunc("hello")
	genericFunc(123)
	fmt.Println(helper.HelperFunc())
}

func genericFunc[k string | int](thing k) {
	fmt.Println(thing)
}
