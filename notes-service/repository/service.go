package repository

type NotesService struct {
	Id        string `dynamo:"id,hash"`
	Name      string `dynamo:"name"`
	Completed bool   `dynamo:"completed"`
}
