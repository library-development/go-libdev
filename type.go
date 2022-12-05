package libdev

type Type struct {
	IsPointer bool   `json:"isPointer"`
	IsArray   bool   `json:"isArray"`
	IsMap     bool   `json:"isMap"`
	BaseType  string `json:"baseType"`
}
