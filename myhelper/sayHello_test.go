package myhelper

import (
	"testing"

	golang_module "github.com/raihanmd/golang-module/v2"
)

func BenchmarkSayHello(b *testing.B){
	for i := 0; i < b.N; i++{
		golang_module.SayHello("Jarem")
	}
}