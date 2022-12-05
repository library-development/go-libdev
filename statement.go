package libdev

type Statement struct {
	Assignments []string `json:"assignments"`
	FunctionID  string   `json:"functionID"`
	Arguments   []string `json:"arguments"`
}
