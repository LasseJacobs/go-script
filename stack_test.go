package main

import "testing"

func TestStack_Peek(t *testing.T) {
	stack := NewStack()
	stack.Push(make(map[string]bool))
	if stack.Peek() == nil {
		t.Fail()
	}
	if stack.Peek() == nil {
		t.Fail()
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack()
	stack.Push(make(map[string]bool))
	stack.Push(make(map[string]bool))
	if stack.Pop() == nil {
		t.Fail()
	}
	if stack.Pop() == nil {
		t.Fail()
	}
	if stack.Pop() != nil {
		t.Fail()
	}
	if !stack.IsEmpty() {
		t.Fail()
	}
}

func TestStack_Get(t *testing.T) {
	stack := NewStack()
	m := make(map[string]bool)
	m["abc"] = true
	stack.Push(make(map[string]bool))
	stack.Push(m)
	stack.Push(make(map[string]bool))
	m2 := stack.Get(1)
	if _, ok := m2["abc"]; ok == false {
		t.Fail()
	}
}
