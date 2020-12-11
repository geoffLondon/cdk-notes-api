package service_resolver

type CallFlowService struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type ServiceParameters struct{}

//todo from schema
// createNote(note: NoteInput!): Note
// createNote(params: CreateNoteParams!): ID! @aws_auth(cognito_groups: ["SmartnumbersAdmin"])

/*defaultAuthorization: {
authorizationType: appSync.AuthorizationType.API_KEY,
apiKeyConfig: {
expires: Expiration.after(Duration.days(365))
}
},*/

/*type Subscription getNotes {
onCreatedNote: Note
@aws_subscribe(mutations: ["createNote"])
}*/

/*input NoteInput {
id: ID!
name: String!
completed: Boolean!
}
*/

// todo cdk-notes-api.GraphQLApiUrl
//  https://x5sgtilnwnearfzmsfbi3k7d54.appsync-api.eu-west-2.amazonaws.com/graphql
