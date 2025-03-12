package main

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestAlecThomasNoContext(t *testing.T) {
	assert.True(t, false)
}

func TestAlecThomasContext(t *testing.T) {
	assert.True(t, false, "foo %v context", "bar")
}

func TestTestifyNoContext(t *testing.T) {
	require.True(t, false)
}

func TestTestifyContext(t *testing.T) {
	require.True(t, false, "foo %v context", "bar")
}
