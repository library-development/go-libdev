package libdev

import (
	"fmt"
	"io"
	"strings"
)

type ImportMap struct {
	names map[string]string
}

func (m *ImportMap) Length() int {
	return len(m.names)
}

func (m *ImportMap) Add(path Path) string {
	name := path.Last().AllLowerNoSpaces()
	for {
		p, ok := m.names[name]
		if !ok {
			break
		}
		if p == path.String() {
			return name
		}
		name = name + "_"
	}
	m.names[name] = path.String()
	return name
}

func (m *ImportMap) Resolve(path Path) string {
	for name, p := range m.names {
		if p == path.String() {
			return name
		}
	}
	return ""
}

func (m *ImportMap) WriteGolang(w io.Writer) error {
	_, err := w.Write([]byte("import (\n"))
	if err != nil {
		return err
	}
	for name, path := range m.names {
		_, err = fmt.Fprintf(w, "\t%s \"%s\"\n", name, strings.TrimPrefix(path, "/"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte(")\n"))
	if err != nil {
		return err
	}
	return nil
}
