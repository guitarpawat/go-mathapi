package preprocess

import (
	"testing"
)

func initTest() {
	memMap = map[string]string{}
}

func Test_Mem_AddToMem_NonNumber(t *testing.T) {
	initTest()
	err := AddToMem("answer", "42.0000")
	if err != nil {
		t.Errorf("error: %s", err)
	}
}

func Test_Mem_AddToMem_Number_ShouldBe_Error(t *testing.T) {
	initTest()
	err := AddToMem("42.0000", "42.0000")
	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_Mem_GetFromMem_AlreadyHas_ShouldBe_Found(t *testing.T) {
	expected := "42.0000"
	initTest()
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
	initTest()
	_, ok := GetFromMem("answer")
	if ok {
		t.Error("ok is true")
	}
}

func Test_RemoveFromMem_NotHas_Ok_ShouldBe_False(t *testing.T) {
	initTest()
	_, ok := RemoveFromMem("answer")
	if ok {
		t.Error("ok is true")
	}
}

func Test_RemoveFromMem_Has_Ok_ShouldBe_False(t *testing.T) {
	expected := "42.0000"
	initTest()
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
	initTest()
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
	initTest()
	err := AddToMem("answer", "par", "20", "add", "1", "end", "mul", "2", "2")
	if err == nil {
		t.Error("expected error")
		return
	}
}
