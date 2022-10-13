package test

import (
	"log"
	"testing"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	log.Println("apakah ini sama", IsSame[string]("yahaha", "yahaha"))
	log.Println("apakah ini sama", IsSame[int](7, 9))
}
