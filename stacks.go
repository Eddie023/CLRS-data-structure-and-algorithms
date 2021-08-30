package main

type Stack struct {
  top int
	value []int
}

// Create a new stack.
func New() *Stack {
	return &Stack{-1, []int{} }
}

// Add element to a stack.
func (s *Stack) Push (elm int) *Stack {
	s.value = append(s.value, elm)
	s.top++
	return s
}

// Pop the top item of the stack and return it.
func (s *Stack) Pop () interface{} {
	if (s.top == -1){
		return nil
	}

	s.top -= 1
	topValue := s.top
	poppedValue := s.value[s.top + 1]

	s.value = append(s.value[:topValue + 1])

	return poppedValue
}