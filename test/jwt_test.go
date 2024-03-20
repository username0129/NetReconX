package test

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"server/internal/utils"
	"testing"
)

func TestJwt(t *testing.T) {
	claims := utils.CustomClaims{
		UUID:        uuid.Must(uuid.NewV4()),
		ID:          0,
		Username:    "admin",
		AuthorityId: 0,
	}
	jwt, err := utils.GenerateJWT(claims)
	if err != nil {
		return
	}
	fmt.Println(jwt)
}
