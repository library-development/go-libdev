package libdev

// Statement is a single statement in a function body.
// Statements can only be in the form `a, b = f(c, d)`.
// FunctionID is the ID of the function being called.
// Inputs is a list of IDs of the variables being passed as inputs to the function.
// Outputs is a list of IDs of the variables being assigned the outputs of the function.
//
// `return` is a special function that ends function execution and returns its inputs.
type Statement struct {
	FunctionID ID
	Inputs     []ID
	Outputs    []ID
}
