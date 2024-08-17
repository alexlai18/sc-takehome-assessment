package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

type Service struct{}

func (s *Service) FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*folders.Folder, error) {
	return folders.FetchAllFoldersByOrgID(orgID)
}

func (s *Service) FetchAllFoldersByOrgIDWithPagination(orgID uuid.UUID, page int, pageSize int) (folders.PaginatedFetchFolderResponse, error) {
	return folders.FetchAllFoldersByOrgIDWithPagination(orgID, page, pageSize)
}

func main() {
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	service := &Service{}
	res, err := folders.GetAllFolders(req, service)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	folders.PrettyPrint(res)
}
