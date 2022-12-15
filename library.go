package libdev

// Library is the top level unit of code distribution.
// Libraries can be compiled into Python, Ruby, Java, C#, TypeScript, Go and Rust.
type Library struct {
	Name       Name
	Version    Version
	Doc        string
	Discussion []Comment

	Constants []Constant
	Types     []Type
	Functions []Function
}
