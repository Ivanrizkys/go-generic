package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestTypeSet(t *testing.T) {
	assert.Equal(t, float32(100), Min[float32](float32(300), float32(100)))
	assert.Equal(t, int32(50), Min[int32](int32(50), int32(600)))
}
