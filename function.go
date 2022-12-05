package libdev

type Function struct {
	Inputs  []Field     `json:"inputs"`
	Outputs []Field     `json:"outputs"`
	Body    []Statement `json:"body"`
}
