package service_resolver

type NotesService struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type ServiceParameters struct{}

/**
TODO
 cdk-notes-api.GraphQLApiUrl
 https://jrtqqn36pzbelfdirfir4ttjny.appsync-api.eu-west-2.amazonaws.com/graphql
*/
