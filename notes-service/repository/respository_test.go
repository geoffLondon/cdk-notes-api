// +build unit

package repository_test

import (
	"context"
	mock_aws_dynamodb "github.com/geoffLondon/cdk-notes-api/mocks/aws/dynamodb"
	notes_service "github.com/geoffLondon/cdk-notes-api/notes-service/repository"
	notes_service_repo "github.com/geoffLondon/cdk-notes-api/notes-service/repository"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestNotesServiceRepository Unit Test", func() {
	ctx := context.TODO()
	var (
		mockController     *gomock.Controller
		mockDynamoDbClient *mock_aws_dynamodb.MockDynamoDbClient
		repo               *notes_service_repo.DynamoServiceRepository
	)

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		mockDynamoDbClient = mock_aws_dynamodb.NewMockDynamoDbClient(mockController)
		repo = notes_service_repo.NewDynamoServiceRepository(mockDynamoDbClient)
	})

	AfterEach(func() {
		mockController.Finish()
	})

	Context("FindAll", func() {
		It("Finds entries when table is not empty", func() {
			expectedOutInterface := make([]notes_service.NotesService, 0)
			mockDynamoDbClient.EXPECT().FindAll(ctx, &expectedOutInterface).Return(nil)

			_, err := repo.FindAll(ctx)

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
