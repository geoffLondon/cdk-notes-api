package service_resolver

type NoteRequest struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type NoteResponse struct {
	Note Note `json:"note"`
}

type Note struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
