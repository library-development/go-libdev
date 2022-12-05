package libdev

import (
	"fmt"
	"io"
)

type Field struct {
	Name Name
	Type Type
}

func (f *Field) WriteGolang(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s ", f.Name.PascalCase())
	if err != nil {
		return err
	}
	return f.Type.WriteGolang(w)
}
