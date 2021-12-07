package pkg_test

import (
	"testing"

	db "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/db/user"
	"github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/jwt"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateJWT(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	if err != nil {
		t.Fatalf("failed to create hash: %v", err)
	}
	uUID, err := uuid.FromString("3cf0316d-a1c4-45dc-b664-7229d89f41d8")
	if err != nil {
		t.Fatalf("failed to create uuid from string: %v", err)
	}
	user := db.User{UUID: uUID, EMAIL: "test@email.com", PERMISSION: "normal", PASSWORD: string(hash)}
	token, err := jwt.CreateJWT(user)
	if err != nil {
		t.Fatalf("failed to create jwt: %v", err)
	}
	assert.Equal(t, token, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiM2NmMDMxNmQtYTFjNC00NWRjLWI2NjQtNzIyOWQ4OWY0MWQ4IiwiUEVSTUlTU0lPTiI6Im5vcm1hbCIsImV4cCI6MTUwMDB9.z_lSE07P7jZzI2eEiq5mU0hojAmwTm5_k6rCCZmUWz8", nil)
}

func TestParseJWT(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiM2NmMDMxNmQtYTFjNC00NWRjLWI2NjQtNzIyOWQ4OWY0MWQ4IiwiUEVSTUlTU0lPTiI6Im5vcm1hbCIsImV4cCI6MTUwMDB9.z_lSE07P7jZzI2eEiq5mU0hojAmwTm5_k6rCCZmUWz8"
	uUID, permission, err := jwt.ParseJWT(token)
	if err != nil {
		t.Fatalf("failed to parse jwt: %v", err)
	}
	assert.Equal(t, uUID, uuid.Must(uuid.FromString("3cf0316d-a1c4-45dc-b664-7229d89f41d8")), nil)
	assert.Equal(t, permission, "normal", nil)
}
