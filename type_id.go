package libdev

type TypeID struct {
	IsMap     bool
	IsArray   bool
	IsPointer bool

	// BaseType is the base type of the type.
	// For example, the base type of "[]string" is "string" and the base type of "map[string]Person" is "Person".
	// The nil value is treated as a string type.
	BaseType ID
}
