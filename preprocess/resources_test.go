package preprocess

import (
	"reflect"
	"testing"
)

func Test_ProcessAll(t *testing.T) {

}

func Test_IsKeyInAnyResource_Constant(t *testing.T) {
	ResetMem()
	AddToMem("answer", "par", "20", "add", "1", "end", "mul", "2")
	actual := ProcessAll([]string{"pi", "answer"})
	expected := []string{"3.141593", "42.0000"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v but get: %v", expected, actual)
	}
}

func Test_IsKeyInAnyResource_Mem(t *testing.T) {
	ResetMem()
	AddToMem("temp", "0")

	expected := true
	actual := IsKeyInAnyResource("temp")
	if expected != actual {
		t.Errorf("expected %t but get: %t", expected, actual)
	}
}

func Test_IsKeyInAnyResource_NotFound(t *testing.T) {
	ResetMem()

	expected := false
	actual := IsKeyInAnyResource("no one")
	if expected != actual {
		t.Errorf("expected %t but get: %t", expected, actual)
	}
}

func Test_IsValidKey_Found(t *testing.T) {
	actual := IsValidKey("mem")
	expected := false

	if expected != actual {
		t.Errorf("expected %t but get: %t", expected, actual)
	}
}

func Test_IsValidKey_NotFound(t *testing.T) {
	actual := IsValidKey("memo")
	expected := true

	if expected != actual {
		t.Errorf("expected %t but get: %t", expected, actual)
	}
}
