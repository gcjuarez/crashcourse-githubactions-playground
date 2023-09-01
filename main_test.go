package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibonacci(t *testing.T) {
	f := fib()
	want, _ := fmt.Println("1 1 2 3 5")
	res, _ := fmt.Println(f(), f(), f(), f(), f())
	require.Equal(t, want, res)
}
