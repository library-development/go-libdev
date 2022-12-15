package libdev

// Function is a reusable unit of execution.
// Functions can call other functions.
// Functions can be called directly from the UI.
type Function struct {
	Name       Name
	Doc        string
	Discussion []Comment

	Inputs []struct {
		Name       Name
		Doc        string
		Discussion []Comment
		Type       TypeID
	}
	Outputs []struct {
		Name       Name
		Doc        string
		Discussion []Comment
		Type       TypeID
	}
	Body []Statement
}
