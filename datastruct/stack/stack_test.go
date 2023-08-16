package stack

import "testing"

func TestListStack(t *testing.T) {
	listStack := NewListStack()

	listStack.Push(1, 1)
	listStack.Push(2, 2)

	t.Logf("%s", listStack.Print())

	for {
		v, has := listStack.Pop()
		if !has {
			break
		}

		t.Logf("pop value %v", v)
	}
}
