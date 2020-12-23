package repository

type NotesService struct {
	//Id        string `dynamo:"id,hash" json:"id"`
	Id        string `dynamo:"id" json:"id"`
	Name      string `dynamo:"name" json:"name"`
	Completed bool   `dynamo:"completed" json:"completed"`
}
