package preprocess

import (
	"testing"
)

func Test_IsKeyInAnyResource_Constant(t *testing.T) {
	expected := true
	actual := IsKeyInAnyResource("phi")
	if expected != actual {
		t.Errorf("expected %t but get: %t", expected, actual)
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
