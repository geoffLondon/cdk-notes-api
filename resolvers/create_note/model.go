package create_note_resolver

type NoteRequest struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}
