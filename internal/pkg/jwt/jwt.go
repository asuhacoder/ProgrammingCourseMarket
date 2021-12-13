package jwt

import (
	"log"
	"time"

	db "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/db/user"
	uuid "github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
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

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

func ParseJWT(tokenString string) (uuid.UUID, string, error) {
	log.Printf("token: %v", tokenString)
	var uUID uuid.UUID
	var permission string
	mySingningKey := []byte("AllYourBase")
	at(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
			return mySingningKey, nil
		})
		log.Println(token)
		if err != nil {
			log.Printf("failed to parse jwt: %v", err)
			return
		} else {
			if claims, ok := token.Claims.(*userClaims); ok && token.Valid {
				uUID = claims.UUID
				permission = claims.PERMISSION
			}
		}
	})
	return uUID, permission, nil
}
