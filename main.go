package main

import (
	"fmt"

	golang_module "github.com/raihanmd/golang-module/v2"
)

func main() {
	hello2 := golang_module.SayGoodbye("Aka")
	hello := golang_module.SayHello("Jarem")
	fmt.Println(hello)
	fmt.Println(hello2)
}