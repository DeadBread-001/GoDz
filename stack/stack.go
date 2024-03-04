package stack

import "errors"

type Stack struct {
	Nodes []any
}

func (s *Stack) Empty() bool {
	return len(s.Nodes) == 0
}

func (s *Stack) Top() (any, error) {
	if s.Empty() {
		return nil, errors.New("tried to get value from empty stack")
	}

	return s.Nodes[len(s.Nodes)-1], nil
}

func (s *Stack) Push(value any) {
	s.Nodes = append(s.Nodes, value)
}

func (s *Stack) Pop() (any, error) {
	if s.Empty() {
		return nil, errors.New("tried to pop from empty stack")
	}

	result := s.Nodes[len(s.Nodes)-1]
	s.Nodes = s.Nodes[:len(s.Nodes)-1]

	return result, nil
}
