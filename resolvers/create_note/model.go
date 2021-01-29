package create_note_resolver

type NoteRequest struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
