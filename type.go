package libdev

import (
	"fmt"
	"io"

	"github.com/function-cafe/go-golang"
)

// Type represents a type in most programming languages.
type Type struct {
	Name       Name
	Doc        string
	Discussion []Comment

	Fields []struct {
		Name       Name
		Doc        string
		Discussion []Comment
		Type       TypeID
	}
}

// WriteGolang writes Go source code to w.
//
// ctx is the path of the package being written to.
// When ctx is empty, the type is always written as a fully qualified type.
func (t *Type) WriteGolang(w io.Writer, ctx Path) error {
	if t == nil {
		return nil
	}

	_, err := w.Write([]byte("package " + ctx.Last().AllLowerNoSpaces() + "\n"))
	if err != nil {
		return err
	}

	importMap := ImportMap{}
	for _, field := range t.Fields {
		importMap.Add(field.Type.BaseType.Path)
	}
	if importMap.Length() > 0 {
		_, err = w.Write([]byte("\n"))
		if err != nil {
			return err
		}
		err = importMap.WriteGolang(w)
		if err != nil {
			return err
		}
	}

	_, err = w.Write([]byte("\n"))
	if err != nil {
		return err
	}

	err = golang.WriteComment(w, t.Doc, 0)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte("type "))
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(t.Name.PascalCase()))
	if err != nil {
		return err
	}

	if len(t.Fields) > 0 {
		_, err = w.Write([]byte(" struct {\n"))
		if err != nil {
			return err
		}

		for _, field := range t.Fields {
			err := golang.WriteComment(w, field.Doc, 1)
			if err != nil {
				return err
			}

			_, err = w.Write([]byte("\t"))
			if err != nil {
				return err
			}

			_, err = w.Write([]byte(field.Name.PascalCase()))
			if err != nil {
				return err
			}

			_, err = w.Write([]byte(" "))
			if err != nil {
				return err
			}

			if field.Type.IsMap {
				_, err := w.Write([]byte("map[string]"))
				if err != nil {
					return err
				}
			}
			if field.Type.IsArray {
				_, err := w.Write([]byte("[]"))
				if err != nil {
					return err
				}
			}
			if field.Type.IsPointer {
				_, err := w.Write([]byte("*"))
				if err != nil {
					return err
				}
			}
			if field.Type.BaseType.Path.Length() == 0 {
				// Default to string type
				_, err := w.Write([]byte("string"))
				return err
			}
			if field.Type.BaseType.Path.Length() == 1 {
				// Handle builtin types
				if !builtinType(field.Type.BaseType.Path.Last().String()) {
					return fmt.Errorf("invalid type: %s", field.Type.BaseType.Path.Last().String())
				}
				_, err := w.Write([]byte(field.Type.BaseType.Path.Last().String()))
				return err
			}
			if !field.Type.BaseType.Path.Pop().Equals(ctx) {
				_, err := w.Write([]byte(importMap.Resolve(field.Type.BaseType.Path.Pop())))
				if err != nil {
					return err
				}
				_, err = w.Write([]byte("."))
				if err != nil {
					return err
				}
			}
			_, err = w.Write([]byte(field.Type.BaseType.Path.Last().PascalCase()))
			if err != nil {
				return err
			}
		}
	}

	_, err = w.Write([]byte("}\n"))
	if err != nil {
		return err
	}

	return nil
}
