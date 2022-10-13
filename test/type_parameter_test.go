package test

import (
	"log"
	"testing"
)

func Length[T any](param T) T {
	return param
}

func TestSample(t *testing.T) {
	str := Length[string]("yahaha")
	number := Length[int](30)

	log.Println(str)
	log.Println(number)
}
