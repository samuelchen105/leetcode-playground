package problem

// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	closeToOpen := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := Stack{}

	for _, r := range s {
		pair, ok := closeToOpen[r]
		if ok {
			if !stack.IsEmpty() && stack.Top() == pair {
				stack.Pop()
			} else {
				return false
			}
		} else {
			stack.Push(r)
		}
	}

	return stack.IsEmpty()
}

type Stack struct {
	inner []rune
}

func (s *Stack) Push(r rune) {
	s.inner = append(s.inner, r)
}

func (s *Stack) Pop() {
	s.inner = s.inner[:len(s.inner)-1]
}

func (s Stack) Top() rune {
	if s.IsEmpty() {
		return 0
	}
	return s.inner[len(s.inner)-1]
}

func (s Stack) IsEmpty() bool {
	return len(s.inner) == 0
}
