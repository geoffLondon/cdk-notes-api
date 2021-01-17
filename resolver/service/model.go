package service_resolver

type NoteRequest struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
