package arrays

import "fmt"

type Stack struct {
	data     []int
	minValue int
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() []int {
	if len(s.data) == 0 {
		fmt.Println("Error: Stack is empty")
		return []int{}
	}
	newData := s.data[:len(s.data)-1]
	s.data = newData
	return newData
}

func (s *Stack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Top() int {
	return s.data[0]
}

func (s *Stack) GetMin() int {
	if len(s.data) == 1 {
		return s.data[0]
	}
	size := len(s.data)
	if s.minValue > s.data[size] && size >= 0 {
		s.minValue = s.data[size]
		size--
		return s.GetMin()
	}

	return s.minValue
}
