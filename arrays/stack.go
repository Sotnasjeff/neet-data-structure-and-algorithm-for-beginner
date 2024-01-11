package arrays

import "fmt"

type Stack struct {
	data []int
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
