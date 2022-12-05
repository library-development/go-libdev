package libdev

import "strings"

type Name struct {
	Words []string
}

func (n *Name) String() string {
	return strings.Join(n.Words, " ")
}

func (n *Name) PascalCase() string {
	for i := range n.Words {
		n.Words[i] = strings.Title(n.Words[i])
	}
	return strings.Join(n.Words, "")
}

func (n *Name) CamelCase() string {
	for i := range n.Words {
		if i == 0 {
			n.Words[i] = strings.ToLower(n.Words[i])
		} else {
			n.Words[i] = strings.Title(n.Words[i])
		}
	}
	return strings.Join(n.Words, "")
}

func (n *Name) SnakeCase() string {
	for i := range n.Words {
		n.Words[i] = strings.ToLower(n.Words[i])
	}
	return strings.Join(n.Words, "_")
}

func (n *Name) KebabCase() string {
	for i := range n.Words {
		n.Words[i] = strings.ToLower(n.Words[i])
	}
	return strings.Join(n.Words, "-")
}

func (n *Name) AllLowerNoSpaces() string {
	return strings.Join(n.Words, "")
}
