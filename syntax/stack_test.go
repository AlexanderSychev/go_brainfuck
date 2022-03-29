package syntax

import (
	"go_brainfuck/runtime"
	"reflect"
	"testing"
)

func TestStack_Length(t *testing.T) {
	t.Log("Length() method test")

	stack := NewStack()
	stack.Push(make([]runtime.Command, 0))

	if stack.Length() != 1 {
		t.Error("Expected 1 as \"Length()\" result")
	}
}

func TestStack_Push(t *testing.T) {
	t.Log("Push() method test")

	stack := NewStack()
	item := make([]runtime.Command, 0)
	stack.Push(item)

	if !reflect.DeepEqual(stack[0], item) {
		t.Error("Expected that same item will be added")
	}
}

func TestStack_Pop(t *testing.T) {
	t.Log("Pop() method tests")

	t.Log("If stack length more then zero then top item will be returned")

	stack01 := NewStack()
	itemToPush := make([]runtime.Command, 0)
	stack01.Push(itemToPush)
	itemToPop := stack01.Pop()

	if !reflect.DeepEqual(itemToPush, itemToPop) {
		t.Error("Expected that same item will popped")
	}

	t.Log("If stack length is zero then empty slice will be returned")

	stack02 := NewStack()
	_ = stack02.Pop()
	if stack02.Length() != 0 {
		t.Error("Item cannot be nil")
	}
}
