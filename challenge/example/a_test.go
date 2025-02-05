package example

import (
	"testing"

	"aoc-go-22/challenge"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := challenge.FromLiteral("42")

	result := a(input)

	require.Equal(t, 42, result)
}
