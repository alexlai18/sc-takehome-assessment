package folders

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started
import (
	"fmt"

	"github.com/gofrs/uuid"
)

type PaginatedFetchFolderResponse struct {
	Folders       []*Folder
	NextPageToken string
}

type PaginatedFetchFolderRequest struct {
	OrgID    uuid.UUID
	Page     int
	PageSize int
}

// COMMENT - all comments for component 1 are in folders.go
// ASSUMPTION - token has been decoded and authorized prior to this function
// The solution is very similar to the explanation, where a page and a page size is passed into the function, which determines which "page" of data is returned
// Account for edge cases where the starting mark exceeds the folders, and if the page + page size overflows the data length as well
// Note: Typically, we would leave pagination to the data retrieval function (or DB itself) - this would be done without the restraints of the task/assessment
// A token is generated, encoded with the page and the pageSize. The pageSize won't change (unless prompted for and changed prior to this function) and the page number is incrememnted into the token
// Token also holds orgID and unique ID. This is for authorization purposes to prevent unknown entities from accessing data (also for logging purposes too)
// If no data left, the token returned will be empty - no more data can be retrieved without extra authorisation
func GetAllFoldersWithPagination(req *PaginatedFetchFolderRequest, service FolderService) (*PaginatedFetchFolderResponse, error) {
	var (
		err error
		fs  PaginatedFetchFolderResponse
	)
	orgID := req.OrgID
	page := req.Page
	pageSize := req.PageSize

	if req.OrgID == uuid.Nil {
		return nil, fmt.Errorf("invalid orgID provided: %s", orgID)
	}

	fs, err = service.FetchAllFoldersByOrgIDWithPagination(orgID, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &fs, err
}

func FetchAllFoldersByOrgIDWithPagination(orgID uuid.UUID, page int, pageSize int) (PaginatedFetchFolderResponse, error) {
	var (
		err           error
		nextPageToken string
	)

	// Typically would have error handling here
	folders := GetSampleData()
	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}

	start := page * pageSize
	end := start + pageSize

	// Handling edge cases for pagination data retrieval
	if start > len(resFolder) {
		return PaginatedFetchFolderResponse{Folders: []*Folder{}, NextPageToken: ""}, nil
	}
	if end > len(resFolder) {
		end = len(resFolder)
		nextPageToken = ""
	} else {
		// If no error, provide them with a token - this token will be parsed at authorization stage
		// Generate NextPageToken to send
		nextPageToken, err = GeneratePaginationToken(page+1, pageSize, orgID, uuid.Must(uuid.NewV4()))
	}

	// Returning paginated pages
	return PaginatedFetchFolderResponse{Folders: resFolder[start:end], NextPageToken: nextPageToken}, err
}
