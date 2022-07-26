package main

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(x map[string]bool) {
	s.list.PushBack(x)
}

func (s *Stack) Pop() map[string]bool {
	if s.list.Len() == 0 {
		return nil
	}
	tail := s.list.Back()
	val := tail.Value
	s.list.Remove(tail)
	return val.(map[string]bool)
}

func (s *Stack) Peek() map[string]bool {
	tail := s.list.Back()
	val := tail.Value
	return val.(map[string]bool)
}

func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) Get(pos int) map[string]bool {
	var it *list.Element = s.list.Front()
	for i := 0; i < pos; i++ {
		it = it.Next()
	}
	return it.Value.(map[string]bool)
}
