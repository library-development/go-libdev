package libdev

import "io"

type Type struct {
	IsPointer bool `json:"isPointer"`
	IsArray   bool `json:"isArray"`
	IsMap     bool `json:"isMap"`
	BaseType  ID   `json:"baseType"`
}

func (t *Type) WriteGolang(w io.Writer) error {
	if t.IsPointer {
		_, err := w.Write([]byte("*"))
		if err != nil {
			return err
		}
	}
	if t.IsArray {
		_, err := w.Write([]byte("[]"))
		if err != nil {
			return err
		}
	}
	if t.IsMap {
		_, err := w.Write([]byte("map[string]"))
		if err != nil {
			return err
		}
	}
	return t.BaseType.WriteGolang(w)
}
