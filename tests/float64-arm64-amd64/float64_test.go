package main

import (
	"fmt"
	"math"
	"testing"
)

// Changing N changes the error, this value has a diff for both amd64 and arm64, but the diff is not the same.
// N is a global to avoid the compiler inlining.
var N = 640002

func TestFloat64Ops(t *testing.T) {
	f := float64(N) / 1e4 * 1e4
	diff := math.Abs(float64(N) - f)
	fmt.Println("N - (N / 1e4 * 1e4) = ", diff)
}
