package libdev

// Constant defines a Constant in a library.
type Constant struct {
	Name       Name
	Doc        string
	Discussion []Comment

	IsExported bool
	Type       TypeID
	Value      any
}
