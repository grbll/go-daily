package validparantheses

type Stack struct {
	Val  rune
	Prev *Stack
}

func isValid(s string) bool {
	var stack *Stack = nil
	for _, paran := range s {
		if paran == '(' || paran == '[' || paran == '{' {
			stack = &Stack{Val: paran, Prev: stack}
		} else {
			if stack == nil || (paran == ')' && stack.Val != '(') || (paran == ']' && stack.Val != '[') || (paran == '}' && stack.Val != '{') {
				return false
			} else {
				stack = stack.Prev
			}
		}
	}
	return stack == nil
}
