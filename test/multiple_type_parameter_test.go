package test

import (
	"fmt"
	"testing"
)

func MultipleTypeParameter[T any, K any](param1 T, param2 K) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleType(t *testing.T) {
	MultipleTypeParameter[string, int]("yagura", 9)
}
