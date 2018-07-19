package preprocess

import (
	"reflect"
	"testing"
)

func Test_MemToNumber(t *testing.T) {
	ResetMem()
	AddToMem("ant", "1.0000")
	AddToMem("answer", "par", "20", "add", "1", "end", "mul", "2")
	actual := MemToNumber([]string{"ant", "answer"})
	expected := []string{"1.0000", "42.0000"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v but get: %v", expected, actual)
	}
}

func Test_Mem_AddToMem_NonNumber(t *testing.T) {
	ResetMem()
	err := AddToMem("answer", "42.0000")
	if err != nil {
		t.Errorf("error: %s", err)
	}
}

func Test_Mem_AddToMem_Number_ShouldBe_Error(t *testing.T) {
	ResetMem()
	err := AddToMem("42.0000", "42.0000")
	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_Mem_GetFromMem_AlreadyHas_ShouldBe_Found(t *testing.T) {
	expected := "42.0000"
	ResetMem()
	_ = AddToMem("answer", expected)
	actual, ok := GetFromMem("answer")
	if !ok {
		t.Error("ok is false")
		return
	}

	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Mem_GetFromMem_AlreadyHas_ShouldBe_NotFound(t *testing.T) {
	ResetMem()
	_, ok := GetFromMem("answer")
	if ok {
		t.Error("ok is true")
	}
}

func Test_RemoveFromMem_NotHas_Ok_ShouldBe_False(t *testing.T) {
	ResetMem()
	_, ok := RemoveFromMem("answer")
	if ok {
		t.Error("ok is true")
	}
}

func Test_RemoveFromMem_Has_Ok_ShouldBe_True(t *testing.T) {
	expected := "42.0000"
	ResetMem()
	_ = AddToMem("answer", expected)

	actual, ok := RemoveFromMem("answer")
	if !ok {
		t.Error("ok is false")
	}

	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Mem_AddToMem_Evaluate(t *testing.T) {
	expected := "42.0000"
	ResetMem()
	err := AddToMem("answer", "par", "20", "add", "1", "end", "mul", "2")
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	actual, ok := GetFromMem("answer")
	if !ok {
		t.Error("ok is false")
		return
	}

	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Mem_AddToMem_Evaluate_ShouldBe_Error(t *testing.T) {
	ResetMem()
	err := AddToMem("answer", "par", "20", "add", "1", "end", "mul", "2", "2")
	if err == nil {
		t.Error("expected error")
	}
}

func Test_Mem_AddToMem_InvalidKey(t *testing.T) {
	ResetMem()
	err := AddToMem("pow", "par", "20", "add", "1", "end", "mul", "2")
	if err == nil {
		t.Error("expected error")
	}
}

func Test_Mem_AddToMem_DuplicateKey(t *testing.T) {
	ResetMem()
	err := AddToMem("e", "par", "20", "add", "1", "end", "mul", "2")
	if err == nil {
		t.Error("expected error")
	}
}

func Test_Mem_AddToMem_UpdateKey_ShouldBe_Allowed(t *testing.T) {
	ResetMem()
	err := AddToMem("memo", "par", "20", "add", "1", "end", "mul", "2")
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	err = AddToMem("memo", "12.3456")
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	expected := "12.3456"
	actual, _ := GetFromMem("memo")

	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}
