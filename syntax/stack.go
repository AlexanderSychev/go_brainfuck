package syntax

import "go_brainfuck/runtime"

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// "Stack" type with methods
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// Stack is implementation of commands slices stack. Need for parsing
type Stack [][]runtime.Command

// Length returns number of stack elements
func (s Stack) Length() int {
	return len(s)
}

// Push adds new top element to stack
func (s *Stack) Push(value []runtime.Command) {
	*s = append(*s, value)
}

// Pop removes and returns top element of stack
func (s *Stack) Pop() []runtime.Command {
	if s.Length() > 0 {
		lastIndex := s.Length() - 1
		lastValue := (*s)[lastIndex]

		*s = (*s)[:lastIndex]

		return lastValue
	} else {
		return make([]runtime.Command, 0)
	}
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// "Stack" constructors
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// NewStack creates empty stack ready to use
func NewStack() Stack {
	return make([][]runtime.Command, 0)
}
