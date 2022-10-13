package test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		log.Println(value)
	}
}

func TestTypeTest(t *testing.T) {
	name := Bag[string]{"lahura", "yahihi", "kuyao", "siti"}
	PrintBag(name)
	assert.Equal(t, true, true)
}
