package stack

import (
	"testing"
)

func Test_Peek_EmptyStack_ShouldBe_Error(t *testing.T) {
	s := NewStack()
	_, err := s.Peek()
	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_Pop_EmptyStack_ShouldBe_Error(t *testing.T) {
	s := NewStack()
	_, err := s.Pop()
	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_Stack_Push_OneItem(t *testing.T) {
	expected := "test"
	s := NewStack()
	s.Push(expected)
	actual, err := s.Peek()

	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}

	if actual != expected {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Stack_Push_TwoItem_ShouldBe_Pop_InOrder(t *testing.T) {
	expected1 := "hello"
	expected2 := "world"
	s := NewStack()
	s.Push(expected2)
	s.Push(expected1)
	actual1, err := s.Pop()

	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}

	if actual1 != expected1 {
		t.Errorf("expected: %s but get: %s", expected1, actual1)
		return
	}

	actual2, err := s.Pop()

	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}

	if actual2 != expected2 {
		t.Errorf("expected: %s but get: %s", expected2, actual2)
	}
}
