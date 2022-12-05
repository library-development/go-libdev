package libdev

type Statement struct {
	Assignments []ID `json:"assignments"`
	FunctionID  ID   `json:"functionID"`
	Arguments   []ID `json:"arguments"`
}
