package jwt

import (
	"time"

	configurations "alshashiguchi/quiz_gem/core"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

//configurations "alshashiguchi/quiz_gem/core"
//SecretKey secret key being used to sign tokens
var (
	SecretKey = []byte(configurations.New().SecretKey.Key)
)

//GenerateToken generates a jwt token
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}

	return tokenString, nil
}

//ParseToken parses a jwt token and returns the user it it's claims
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}

}
