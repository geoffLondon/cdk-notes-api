// +build unit

package create_note_resolver_test

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/lambdacontext"
	create_note_fixtures "github.com/geoffLondon/cdk-notes-api/fixtures/create_note"
	mock_notes_service_repository "github.com/geoffLondon/cdk-notes-api/mocks/notes-service/repository"
	mock_create_note_validator "github.com/geoffLondon/cdk-notes-api/mocks/resolvers/create_note"
	mock_uuid_generator "github.com/geoffLondon/cdk-notes-api/mocks/uuid"
	service_repository "github.com/geoffLondon/cdk-notes-api/notes-service/repository"
	create_note_resolver "github.com/geoffLondon/cdk-notes-api/resolvers/create_note"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateNoteResolver Test", func() {
	ctx := context.TODO()
	lc := new(lambdacontext.LambdaContext)
	ctx = lambdacontext.NewContext(ctx, lc)

	var (
		controller *gomock.Controller
		validator  *mock_create_note_validator.MockValidator
		repository *mock_notes_service_repository.MockServiceRepository
		uuid       *mock_uuid_generator.MockUuidGenerator

		resolver create_note_resolver.CreateNoteResolver
	)

	BeforeEach(func() {
		controller = gomock.NewController(GinkgoT())
		defer controller.Finish()

		validator = mock_create_note_validator.NewMockValidator(controller)
		repository = mock_notes_service_repository.NewMockServiceRepository(controller)
		uuid = mock_uuid_generator.NewMockUuidGenerator(controller)

		resolver = create_note_resolver.NewDefaultCreateNoteResolver(repository, validator, uuid)

	})

	Describe(".Handle()", func() {
		var (
			request create_note_resolver.NoteRequest
			service service_repository.NotesService
		)

		BeforeEach(func() {
			request = create_note_fixtures.ValidCreateNoteRequest()
			service = service_repository.NotesService{
				Text: "lets make notes",
				Done: true,
			}
		})

		Context("Resolver Handler success", func() {
			It("Returns true", func() {
				uuid.EXPECT().New()
				validator.EXPECT().Validate(request).Return(nil)
				repository.EXPECT().Save(ctx, service).Return(nil)

				result, err := resolver.Handle(ctx, request)

				Expect(result).To(BeTrue())
				Expect(err).To(BeNil())

			})
		})

		Context("Resolver Handler failure", func() {
			It("Returns false", func() {
				errFindId := errors.New("text missing")
				uuid.EXPECT().New()
				validator.EXPECT().Validate(request).Return(nil)
				repository.EXPECT().Save(ctx, service).Return(errFindId)

				result, err := resolver.Handle(ctx, request)

				Expect(result).To(BeFalse())
				Expect(err).To(MatchError("text missing"))

			})
		})

	})
})
