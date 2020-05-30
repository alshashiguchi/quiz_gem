package jwt

import (
	"time"

	"alshashiguchi/quiz_gem/graph/model"
	"alshashiguchi/quiz_gem/models/users"

	configurations "alshashiguchi/quiz_gem/core"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

//SecretKey secret key being used to sign tokens
var (
	SecretKey = []byte(configurations.New().SecretKey.Key)
)

//GenerateToken generates a jwt token
func GenerateToken(user model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = user.Username
	claims["access"] = user.Access.String()
	claims["situation"] = user.Situation.String()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}

	return tokenString, nil
}

//ParseToken parses a jwt token and returns the user it it's claims
func ParseToken(tokenStr string) (users.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user users.User

		user.Username = claims["username"].(string)
		user.Access = returnAccess(claims["access"].(string))
		user.Situation = returnSituation(claims["username"].(string))

		return user, nil
	} else {
		return users.User{}, err
	}

}

func returnSituation(situation string) model.UserStatus {
	switch situation {
	case "ACTIVE":
		return model.UserStatusActive
	case "INACTIVE":
		return model.UserStatusInactive
	default:
		return model.UserStatusBlocked
	}

}

func returnAccess(access string) model.Access {
	switch access {
	case "STUDENT":
		return model.AccessStudent
	case "INSTRUCTOR":
		return model.AccessInstructor
	case "ADMIN":
		return model.AccessAdmin
	default:
		return model.AccessNoaccess
	}
}
