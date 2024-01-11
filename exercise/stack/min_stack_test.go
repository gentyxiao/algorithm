package stack

import "testing"

func TestMinStack(t *testing.T) {
	minStack := ContructorMinStack()
	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(3)
	minStack.Push(4)
	got := minStack.GetMin()
	want := 1
	if got == want {
		t.Logf("test min stack success")
	} else {
		t.Errorf("test min stacj fail")
	}

}
