package condition_tree

type stack []string

// IsEmpty check if stack is empty
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Pop Remove and return top element of stack. Return false if stack is empty.
func (s *stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *stack) contains(ele string) bool {
	for _, str := range *s {
		if str == ele {
			return true
		}
	}
	return false
}
