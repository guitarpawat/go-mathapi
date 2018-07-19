package preprocess

import (
	"fmt"

	"github.com/guitarpawat/go-mathapi/evaluate"
)

var memMap = map[string]string{}

// MemToNumber replaces mem specify in mem map.
func MemToNumber(in []string) []string {
	for i := 0; i < len(in); i++ {
		con, ok := GetFromMem(in[i])
		if ok {
			in[i] = con
		}
	}
	return in
}

// AddToMem stores memory(value) to the key of map
func AddToMem(key string, infix ...string) error {
	if evaluate.IsNumber(key) {
		return fmt.Errorf("%s is a number", key)
	}

	_, ok := GetFromMem(key)
	if (IsKeyInAnyResource(key) && !ok) || !IsValidKey(key) {
		return fmt.Errorf("cannot remember using name: %s", key)
	}

	result, err := evaluate.Evaluate(infix)
	if err != nil {
		return fmt.Errorf("cannot add to memory: %s", err)
	}

	memMap[key] = result
	return nil
}

// GetFromMem gets value of the specified key in the map
func GetFromMem(key string) (string, bool) {
	value, ok := memMap[key]
	return value, ok
}

// RemoveFromMem removes value of the specified key from the map
func RemoveFromMem(key string) (string, bool) {
	value, ok := GetFromMem(key)
	if !ok {
		return "", false
	}

	delete(memMap, key)
	return value, true
}

// ResetMem clears the map
func ResetMem() {
	memMap = map[string]string{}
}
