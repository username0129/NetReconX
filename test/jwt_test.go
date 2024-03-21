package test

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"testing"
)

func TestJwt(t *testing.T) {
	claims := util.CustomClaims{
		UUID:        uuid.Must(uuid.NewV4()),
		ID:          0,
		Username:    "admin",
		AuthorityId: 0,
	}
	jwt, err := util.GenerateJWT(claims)
	if err != nil {
		return
	}
	fmt.Println(jwt)
}
