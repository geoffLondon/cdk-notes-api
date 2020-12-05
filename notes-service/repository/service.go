package repository

type NotesService struct {
	Id        string `dynamo:"pk,hash"`
	Name      string `dynamo:"name"`
	Completed bool   `dynamo:"dn"`
}
