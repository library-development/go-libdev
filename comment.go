package libdev

type Comment struct {
	Timestamp int64     `json:"timestamp"`
	AuthorID  string    `json:"author_id"`
	Text      string    `json:"text"`
	Replies   []Comment `json:"replies"`
}
