package libdev

import (
	"fmt"
	"html/template"
	"io"
)

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

// WriteHTML writes the library to the given writer as HTML.
func (l *Library) WriteHTML(w io.Writer) error {
	template, err := template.New("library").Parse(libraryTemplate)
	if err != nil {
		return err
	}
	return template.Execute(w, l)
}

func (l *Library) WritePython(dir string) error {
	return fmt.Errorf("not implemented")
}

func (l *Library) WriteRuby(dir string) error {
	return fmt.Errorf("not implemented")
}

func (l *Library) WriteJava(dir string) error {
	return fmt.Errorf("not implemented")
}

func (l *Library) WriteCSharp(dir string) error {
	return fmt.Errorf("not implemented")
}

func (l *Library) WriteTypeScript(dir string) error {
	return fmt.Errorf("not implemented")
}

func (l *Library) WriteGo(dir string) error {
	return fmt.Errorf("not implemented")
}

func (l *Library) WriteRust(dir string) error {
	return fmt.Errorf("not implemented")
}
