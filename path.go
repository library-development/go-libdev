package libdev

import "strings"

type Path []Name

func (p Path) String() string {
	var s strings.Builder
	for _, n := range p {
		s.WriteRune('/')
		s.WriteString(n.AllLowerNoSpaces())
	}
	return s.String()
}

func (p Path) Length() int {
	return len(p)
}

func (p Path) Last() *Name {
	return &p[len(p)-1]
}

func (p Path) SecondLast() *Name {
	return &p[len(p)-2]
}

func (p Path) Pop() Path {
	return p[:len(p)-1]
}

func (p Path) Equals(other Path) bool {
	if len(p) != len(other) {
		return false
	}
	for i := range p {
		if p[i].AllLowerNoSpaces() != other[i].AllLowerNoSpaces() {
			return false
		}
	}
	return true
}
