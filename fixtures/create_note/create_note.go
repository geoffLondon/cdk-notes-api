// +build unit

package create_note_fixtures

import (
	create_note_resolver "github.com/geoffLondon/cdk-notes-api/resolvers/create_note"
)

func ValidCreateNoteRequest() create_note_resolver.NoteRequest {
	return create_note_resolver.NoteRequest{
		Text: "lets make notes",
		Done: true,
	}
}

func InValidCreateNoteRequest() create_note_resolver.NoteRequest {
	return create_note_resolver.NoteRequest{
		Text: "",
		Done: false,
	}
}
