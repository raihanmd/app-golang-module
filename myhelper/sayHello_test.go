package myhelper 

import "testing"

func TestSayHelloWorld(t *testing.T){
	result := SayHelloWorld("Aku")
	if result != "Aku Said, Hello World!!"{
		panic("Result is not Aku Said, Hello World!!")
	}
}

func TestSayHelloWorldAgus(t *testing.T){
	result := SayHelloWorld("Agus")
	if result != "Agus Said, Hello World!!"{
		panic("Result is not Agus Said, Hello World!!")
	}
}