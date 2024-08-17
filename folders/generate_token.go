package folders

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

// TokenData holds the information to be encoded in the token
type TokenData struct {
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
	OrgID    uuid.UUID `json:"org_id"`
	UUID     uuid.UUID `json:"id"`
	jwt.StandardClaims
}

// Would typically be stored in env
var jwtSecret = []byte("random_secret")

// Create token for next page - should include page, pageSize, orgID, and user info (will randomly make uuid)
func GeneratePaginationToken(page int, pageSize int, orgID uuid.UUID, id uuid.UUID) (string, error) {
	data := TokenData{
		Page:     page,
		PageSize: pageSize,
		OrgID:    orgID,
		UUID:     id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
