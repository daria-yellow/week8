package brackets

func Bracket(s string) (bool, error) {
	stack := New()
	for _, val := range s {
		if (val == ')' || val == '}' || val == ']') && len(stack.I) == 0 {
			return false, nil
		}
		if val == '(' || val == '{' || val == '[' {
			stack.Push(int(val))
		}
		if val == ')' || val == '}' || val == ']' {
			if int(val)-stack.Pop() > 2 {
				return false, nil
			}
		}
	}

	if len(stack.I) != 0 {
		return false, nil
	}

	return true, nil
}

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
