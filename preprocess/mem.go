package preprocess

import (
	"fmt"

	"github.com/guitarpawat/go-mathapi/evaluate"
)

var memMap = map[string]string{}

// AddToMem stores memory(value) to the key of map
func AddToMem(key string, infix ...string) error {
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
