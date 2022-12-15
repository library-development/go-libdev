package libdev

// BuiltinType returns true if s is a builtin type.
func builtinType(s string) bool {
	if s == "string" {
		return true
	}
	if s == "byte" {
		return true
	}
	if s == "bool" {
		return true
	}
	if s == "int" {
		return true
	}
	// TODO: fill out the rest of the builtin types here.
	return false
}
