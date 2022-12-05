package libdev

import (
	"fmt"
	"io"
)

type ID struct {
	Path []Name
}

func (id *ID) WriteGolang(w io.Writer) error {
	l := len(id.Path)
	if l == 0 {
		return nil
	} else if l == 1 {
		_, err := fmt.Fprintf(w, "%s", id.Path[0].CamelCase())
		if err != nil {
			return err
		}
	} else {
		_, err := fmt.Fprintf(w, "%s.%s", id.Path[l-2].CamelCase(), id.Path[l-1].PascalCase())
		if err != nil {
			return err
		}
	}
	return nil
}
