package valid_parentheses

type stack[T any] struct {
	Push   func(T)
	Pop    func() T
	Length func() int
}

func Stack[T any]() stack[T] {
	slice := make([]T, 0)
	return stack[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Pop: func() T {
			res := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			return res
		},
		Length: func() int {
			return len(slice)
		},
	}
}

func IsValid(s string) bool {
	stack := Stack[string]()
	p := make(map[string]string)
	p[")"] = "("
	p["]"] = "["
	p["}"] = "{"
	for _, r := range s {
		c := string(r)
		switch c {
		case "(", "[", "{":
			stack.Push(c)
		case ")", "]", "}":
			poppedV := stack.Pop()
			if poppedV != p[c] {
				return false
			}
		}
	}
	return stack.Length() == 0
}
