package preprocess

// IsKeyInAnyResource returns true if any map has the key.
func IsKeyInAnyResource(key string) bool {
	_, ok1 := GetConstant(key)
	if ok1 {
		return ok1
	}
	_, ok2 := GetFromMem(key)
	if ok2 {
		return ok2
	}

	return false
}
