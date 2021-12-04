package jwt

import (
	db "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/db/user"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

type userClaims struct {
	UUID       uuid.UUID
	PERMISSION string
	jwt.StandardClaims
}

func CreateJWT(user db.User) (string, error) {
	mySingningKey := []byte("AllYourBase")

	claims := userClaims{
		user.UUID,
		user.PERMISSION,
		jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySingningKey)

	return ss, err
}
