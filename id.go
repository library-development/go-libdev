package libdev

import (
	"fmt"
	"io"
)

type ID struct {
	Path Path
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

func (id *ID) WriteRuby(w io.Writer, kase string) error {
	for i := range id.Path {
		if i == len(id.Path)-1 {
			if i > 0 {
				_, err := w.Write([]byte("."))
				if err != nil {
					return err
				}
			}
			_, err := fmt.Fprintf(w, "%s", id.Path[i].Case(kase))
			if err != nil {
				return err
			}
		} else {
			if i > 0 {
				_, err := w.Write([]byte("::"))
				if err != nil {
					return err
				}
			}
			_, err := fmt.Fprintf(w, "%s", id.Path[i].CamelCase())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (id *ID) WritePython(w io.Writer, kase string) error {
	l := len(id.Path)
	if l == 0 {
		return nil
	} else if l == 1 {
		if id.Path[0].String() == "string" {
			_, err := w.Write([]byte("str"))
			return err
		}
	}
	_, err := fmt.Fprintf(w, "%s", id.Path[l-1].Case(kase))
	return err
}

func (id *ID) WriteTypescript(w io.Writer, kase string) error {
	l := len(id.Path)
	if l == 0 {
		return nil
	} else if l == 1 {
		if id.Path[0].String() == "string" {
			_, err := w.Write([]byte("str"))
			return err
		}
	}
	_, err := fmt.Fprintf(w, "%s", id.Path[l-1].Case(kase))
	return err
}
