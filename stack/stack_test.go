package stack_test

import (
	"testing"

	"github.com/DeadBread-001/GoDz/tree/dz1part2/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Parallel()

	var s stack.Stack
	assert.Equal(t, 0, len(s.Nodes))

	s.Push("first")
	assert.Equal(t, 1, len(s.Nodes))

	top, _ := s.Top()
	assert.Equal(t, "first", top)
	assert.Equal(t, 1, len(s.Nodes))

	popped, _ := s.Pop()
	assert.Equal(t, "first", popped)
	assert.Equal(t, 0, len(s.Nodes))

	s.Push("1")
	s.Push("2")
	s.Push("3")
	s.Push("4")
	assert.Equal(t, 4, len(s.Nodes))
	popped, _ = s.Pop()
	assert.Equal(t, "4", popped)
	assert.Equal(t, 3, len(s.Nodes))
}

func TestStackErrors(t *testing.T) {
	t.Parallel()

	var s stack.Stack

	_, err := s.Top()
	assert.NotNil(t, err)

	_, err = s.Pop()
	assert.NotNil(t, err)
}
