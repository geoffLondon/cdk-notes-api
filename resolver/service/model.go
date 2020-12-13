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
 https://exzqcxlq75dt3appoeg2266gbi.appsync-api.eu-west-2.amazonaws.com/graphql
*/
