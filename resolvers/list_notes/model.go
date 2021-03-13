package list_notes_resolver

type ListNotesResponse struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}
