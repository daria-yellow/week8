package stack

type Stack struct {
	I []int
}

func New() *Stack {
	i := make([]int, 0)
	return &Stack{i}
}

func (s *Stack) Push(i int) {
	s.I = append(s.I, i)
}

func (s *Stack) Pop() int {
	x := s.I[len(s.I)-1]
	s.I = s.I[:len(s.I)-1]
	return x
}
