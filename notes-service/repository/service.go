package repository

type NotesService struct {
	Id   string `dynamo:"id,hash" json:"id"`
	Text string `dynamo:"text" json:"text"`
	Done bool   `dynamo:"done" json:"done"`
}
