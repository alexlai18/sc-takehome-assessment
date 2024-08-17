package folders

import (
	"fmt" // Added mew import for error handling

	"github.com/gofrs/uuid"
)

type FolderService interface {
	FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error)
	FetchAllFoldersByOrgIDWithPagination(orgID uuid.UUID, page int, pageSize int) (PaginatedFetchFolderResponse, error)
}

// COMMENT - we should utilise predefined variables rather than intialising new ones
// COMMENT - more error handling is to be done (use err variable)
// COMMENT - no reason to convert from []*Folder to []Folder back to []*Folder
func GetAllFolders(req *FetchFolderRequest, service FolderService) (*FetchFolderResponse, error) {
	var (
		err error
		// f1  Folder - IMPROVEMENT/FIX - removed as we often work with the address of folders rather than the actual folder
		fs []*Folder
	)

	// IMPROVEMENT - Check if orgID is nil
	if req.OrgID == uuid.Nil {
		return nil, fmt.Errorf("invalid orgID provided: %s", req.OrgID)
	}

	// LOGIC - Fetches all folders (in the form of array of pointers to Folders) that belong to a specific organisation given the ID (and returns error signal)
	fs, err = service.FetchAllFoldersByOrgID(req.OrgID)
	// IMPROVEMENT - Handling any errors returns from FetchAllFoldersByOrgId function
	if err != nil {
		return nil, err
	}

	// BELOW INCLUDES IMPROVEMENTS TO CODE AND COMPILER ERRORS. THIS WILL NOT BE USED AS WE WILL BE SIMPLIFYING LOGIC FLOW BY AVOIDING UNECESSARY CONVERSIONS
	/*
			   // LOGIC - Iterates through folders and dereferences the Folder pointer and appends it to list of folders
			   f := []Folder{}        // IMPROVEMENT - Moved this declaration as it is only used from this point onwards
			   for _, v := range fs { // FIX - Has compiler error - bypass by declaring an unused variable
			       f = append(f, *v)
			   }

			   var fp []*Folder
			   // LOGIC - Appends address of the folder to fp folder array
			   // FIX - THIS HAS A COMMON PITFALL - since v1 is reused, appending its address to the slice will point to the last Folder that v1 is declared as (last folder)
			   for k1 := range f { // FIX - Removed v1 to bypass address discrepency with iterators - using slice indexing instead
			       fp = append(fp, &f[k1])
			   }

			   // IMRPVOEMENT - Copy all data into the response object - rather than doing this, we can make more concise by directly returning the response object
			   // var ffr *FetchFolderResponse
			   // ffr := &FetchFolderResponse{Folders: fp}
		       // return &FetchFolderResponse{Folders: fp}, err
	*/
	return &FetchFolderResponse{Folders: fs}, err
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	// LOGIC - This gets the folder data (in this case is a JSON but in practice will be from DB) - in form of []*Folder
	// NOTE - I won't change GetSampleData function - but it should also include an error return value rather than panic so we can handle
	folders := GetSampleData()

	resFolder := []*Folder{}
	// LOGIC - Iterates through folders - if the OrgID is same as the param - append to resFolder to return
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	// Returns slice of folder pointers AND nil which signifies that there is no error
	return resFolder, nil
}
