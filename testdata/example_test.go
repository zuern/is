package example

// CAUTION: DO NOT EDIT
// Tests in this project rely on specific lines numbers
// throughout this file.

import (
	"math"
	"testing"

	"github.com/matryer/is"
)

func TestSomething(t *testing.T) {
	// this comment will be extracted
}

func TestSomethingElse(t *testing.T) {
	is := is.New(t)
	a, b := 1, 2
	getB := func() int {
		return b
	}
	is.True(a == getB()) // should be the same
}

func TestSomethingElseTpp(t *testing.T) {
	is := is.New(t)
	a, b := 1, 2
	getB := func() int {
		return b
	}
	is.True(a == getB())
}

func TestMinVal(t *testing.T) {
	is := is.New(t)
	minVal := func(a, b float64) float64 {
		return math.Min(a, b)
	}
	is.True(minVal(1, 2) == 1, "should be the same")
}
