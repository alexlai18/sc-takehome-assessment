package folders_test

import (
	"fmt"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// NOTE - BIG TESTING INSPIRATION FROM https://dev.to/salesforceeng/mocks-in-go-tests-with-testify-mock-6pd AND https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/

type MockFolderService struct {
	mock.Mock
}

func (mock *MockFolderService) FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*folders.Folder, error) {
	args := mock.Called(orgID)
	return args.Get(0).([]*folders.Folder), args.Error(1)
}

func (m *MockFolderService) FetchAllFoldersByOrgIDWithPagination(orgID uuid.UUID, page int, pageSize int) (folders.PaginatedFetchFolderResponse, error) {
	args := m.Called(orgID, page, pageSize)
	return args.Get(0).(folders.PaginatedFetchFolderResponse), args.Error(1)
}

func Test_GetAllFolders(t *testing.T) {
	mockService := new(MockFolderService)
	testingOrgID := uuid.Must(uuid.NewV4())
	folder1 := &folders.Folder{OrgId: testingOrgID}
	folder2 := &folders.Folder{OrgId: testingOrgID}

	// Add your tests here
	var tests = []struct {
		name     string
		req      *folders.FetchFolderRequest
		expected *folders.FetchFolderResponse
		err      error
	}{
		{
			name:     "Valid OrgID returns corresponding folders",
			req:      &folders.FetchFolderRequest{OrgID: testingOrgID},
			expected: &folders.FetchFolderResponse{Folders: []*folders.Folder{folder1, folder2}},
			err:      nil,
		},
		{
			name:     "Invalid OrgID",
			req:      &folders.FetchFolderRequest{OrgID: uuid.Nil},
			expected: nil,
			err:      fmt.Errorf("invalid orgID provided: %s", uuid.Nil),
		},
		{
			name:     "Valid OrgID with no folders",
			req:      &folders.FetchFolderRequest{OrgID: uuid.Must(uuid.NewV4())},
			expected: &folders.FetchFolderResponse{Folders: []*folders.Folder{}}, // Empty slice
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test
			if tt.req.OrgID != uuid.Nil {
				mockService.On("FetchAllFoldersByOrgID", tt.req.OrgID).Return(tt.expected.Folders, tt.err)
			}
			result, err := folders.GetAllFolders(tt.req, mockService)

			if err != nil {
				assert.Equal(t, tt.err, err, "Expected error did not match")
			} else {
				assert.Equal(t, tt.expected, result, "Expected result did not match")
			}
			mockService.AssertExpectations(t)
		})
	}
}
