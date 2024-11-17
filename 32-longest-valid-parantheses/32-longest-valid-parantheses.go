package longestvalidparantheses

type node struct {
	preCurrent int
	next       *node
}

type stack struct {
	top *node
}

func (k *stack) push(current int) {
	if k.top == nil {
		k.top = &node{preCurrent: current, next: nil}
	} else {
		k.top = &node{preCurrent: current, next: k.top}
	}
}

func (k *stack) pop() (int, bool) {
	if k.top == nil {
		return 0, false
	} else {
		current := k.top.preCurrent
		k.top = k.top.next
		return current, true
	}
}

func longestValidParentheses(s string) int {
	var maxi, current int = 0, 0
	var k stack
	for _, char := range s {
		if char == '(' {
			k.push(current)
			current = 0
		} else if preCurrent, ok := k.pop(); ok {
			current += preCurrent + 2
			maxi = max(maxi, current)
		} else {
			current = 0
		}
	}
	return maxi
}
